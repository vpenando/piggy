package main

import (
	"encoding/json"
	"fmt"
	"localization"
	"log"
	"net/http"
	"routing"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	homePageTemplate     localization.HomePageTemplate
	editPageTemplate     localization.EditPageTemplate
	settingsPageTemplate localization.SettingsPageTemplate
	database             *gorm.DB
	// operationController  *piggy.OperationController
	// categoryController   *piggy.CategoryController
)

func init() {
	var err error
	database, err = gorm.Open(sqlite.Open(serverDatabase), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to init database: %s", err))
	}
	// operationController, _ = piggy.NewOperationController(database)
	// categoryController, _ = piggy.NewCategoryController(database)
	routing.InitControllers(database)
}

func handleRoutes() {
	r := mux.NewRouter()
	// Pages
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/edit", edit).Methods("GET").Queries(
		"year", "{year}", // ?year=...
		"month", "{month}", // &month=...
	)
	r.HandleFunc("/settings", settings).Methods("GET")

	// Resources
	r.HandleFunc("/css/{sheet}", styles).Methods("GET")
	r.HandleFunc("/images/{img}", images).Methods("GET")
	r.HandleFunc("/scripts/{script}", scripts).Methods("GET")

	// Misc
	r.HandleFunc("/months", months).Methods("GET")
	r.HandleFunc("/reports", reports).Methods("GET").Queries(
		"year", "{year}", // ?year=...
		"month", "{month}", // &month=...
	)

	// Operations
	r.HandleFunc("/operations", routing.GetOperations).Methods("GET").Queries(
		"year", "{year}", // ?year=...
		"month", "{month}", // &month=...
	)
	r.HandleFunc("/operations", routing.PostOperations).Methods("POST")
	r.HandleFunc("/operations", routing.UpdateOperations).Methods("PUT")
	r.HandleFunc("/operations", routing.DeleteOperations).Methods("DELETE")

	// Categories
	r.HandleFunc("/categories", routing.GetCategories).Methods("GET")
	r.HandleFunc("/categories/{img}", getCategoryIcon).Methods("GET")
	r.HandleFunc("/categories", routing.PostCategory).Methods("POST")

	port := serverPort
	srv := &http.Server{
		Handler: r,
		Addr:    ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on port", port)
	log.Fatal(srv.ListenAndServe())
}

const (
	homeTemplate     = "./views/home.html"
	editTemplate     = "./views/edit.html"
	settingsTemplate = "./views/settings.html"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	t := template.Must(template.ParseFiles(homeTemplate))
	year := time.Now().Year()
	month := time.Now().Month()
	var err error
	homePageTemplate, err = localization.NewHomePageTemplate(year, month, currentLanguage)
	if err != nil {
		routing.HandleError(w, err, http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "home", homePageTemplate)
	if err != nil {
		routing.HandleError(w, err, http.StatusInternalServerError)
		return
	}
}

func edit(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	t := template.Must(template.ParseFiles(editTemplate))
	vars := mux.Vars(r)
	year, err := routing.ParseVarYear(vars)
	if err != nil {
		routing.HandleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	month, err := routing.ParseVarMonth(vars)
	if err != nil {
		routing.HandleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	editPageTemplate, err = localization.NewEditPageTemplate(year, month, currentLanguage)
	if err != nil {
		routing.HandleError(w, err, http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "edit", editPageTemplate)
	if err != nil {
		routing.HandleError(w, err, http.StatusInternalServerError)
		return
	}
}

func settings(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	t := template.Must(template.ParseFiles(settingsTemplate))
	var err error
	settingsPageTemplate, err = localization.NewSettingsPageTemplate(currentLanguage, serverPort, serverDatabase)
	if err != nil {
		routing.HandleError(w, err, http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "settings", settingsPageTemplate)
	if err != nil {
		routing.HandleError(w, err, http.StatusInternalServerError)
		return
	}
}

func getCategoryIcon(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	vars := mux.Vars(r)
	img := "./categories/" + vars["img"]
	routing.ServeImage(w, r, img)
}

func images(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	vars := mux.Vars(r)
	img := "./images/" + vars["img"]
	routing.ServeImage(w, r, img)
}

func scripts(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	r.Header.Set("Content-Type", "application/javascript")
	vars := mux.Vars(r)
	script := "./scripts/" + vars["script"]
	http.ServeFile(w, r, script)
}

func styles(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	r.Header.Set("Content-Type", "text/css")
	vars := mux.Vars(r)
	sheet := "./css/" + vars["sheet"]
	http.ServeFile(w, r, sheet)
}

func months(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	months := localization.MonthsByLanguage(currentLanguage)
	serialized, _ := json.Marshal(months)
	w.WriteHeader(http.StatusOK)
	w.Write(serialized)
}

const reportFilename = "./reports/report.xlsx"

func reports(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	vars := mux.Vars(r)
	year, err := routing.ParseVarYear(vars)
	if err != nil {
		routing.HandleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	month, err := routing.ParseVarMonth(vars)
	if err != nil {
		routing.HandleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	file, err := routing.NewReport(year, month)
	if err != nil {
		routing.HandleError(w, err, http.StatusInternalServerError)
		return
	}
	err = routing.Export(reportFilename, file, currentLanguage)
	if err != nil {
		routing.HandleError(w, err, http.StatusInternalServerError)
		return
	}
	r.Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	http.ServeFile(w, r, reportFilename)
}
