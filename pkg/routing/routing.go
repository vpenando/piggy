package routing

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/vpenando/piggy/pkg/config"
	"github.com/vpenando/piggy/pkg/localization"
	"github.com/vpenando/piggy/pkg/piggy"
)

var (
	homePageTemplate     localization.HomePageTemplate
	editPageTemplate     localization.EditPageTemplate
	settingsPageTemplate localization.SettingsPageTemplate
)

const (
	homeTemplate     = "./static/views/home.html"
	editTemplate     = "./static/views/edit.html"
	settingsTemplate = "./static/views/settings.html"

	reportFilename = "./reports/report.xlsx"
)

var (
	operationController piggy.OperationController
	categoryController  piggy.CategoryController
)

var (
	database *gorm.DB
)

func init() {
	var err error
	database, err = gorm.Open(sqlite.Open(config.ServerDatabase), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to init database: %s", err))
	}
	operationController, _ = piggy.NewOperationController(database)
	categoryController, _ = piggy.NewCategoryController(database)
}

// HandleRoutes starts listening.
func HandleRoutes() {
	r := mux.NewRouter()
	// Pages
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/edit", edit).Methods("GET").Queries(
		"year", "{year}", // ?year=...
		"month", "{month}", // &month=...
	)
	r.HandleFunc("/settings", settings).Methods("GET")

	// Resources
	r.HandleFunc("/static/css/{sheet}", styles).Methods("GET")
	r.HandleFunc("/static/images/{img}", images).Methods("GET")
	r.HandleFunc("/static/scripts/{script}", scripts).Methods("GET")

	// Misc
	r.HandleFunc("/months", months).Methods("GET")
	r.HandleFunc("/reports", reports).Methods("GET").Queries(
		"year", "{year}", // ?year=...
		"month", "{month}", // &month=...
	)

	// Operations
	r.HandleFunc("/operations", getOperations).Methods("GET").Queries(
		"year", "{year}", // ?year=...
		"month", "{month}", // &month=...
	)
	r.HandleFunc("/operations", postOperations).Methods("POST")
	r.HandleFunc("/operations", updateOperations).Methods("PUT")
	r.HandleFunc("/operations", deleteOperations).Methods("DELETE")

	// Categories
	r.HandleFunc("/categories", getCategories).Methods("GET")
	r.HandleFunc("/categories/{img}", getCategoryIcon).Methods("GET")
	r.HandleFunc("/categories", postCategory).Methods("POST")

	port := config.ServerPort
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

func home(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	t := template.Must(template.ParseFiles(homeTemplate))
	year := time.Now().Year()
	month := time.Now().Month()
	var err error
	homePageTemplate, err = localization.NewHomePageTemplate(year, month, config.CurrentLanguage)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "home", homePageTemplate)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}

func edit(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	t := template.Must(template.ParseFiles(editTemplate))
	vars := mux.Vars(r)
	year, err := parseVarYear(vars)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	month, err := parseVarMonth(vars)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	editPageTemplate, err = localization.NewEditPageTemplate(year, month, config.CurrentLanguage)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "edit", editPageTemplate)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}

func settings(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	t := template.Must(template.ParseFiles(settingsTemplate))
	var err error
	settingsPageTemplate, err = localization.NewSettingsPageTemplate(config.CurrentLanguage, config.ServerPort)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "settings", settingsPageTemplate)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}

func getCategoryIcon(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	vars := mux.Vars(r)
	img := "./categories/" + vars["img"]
	serveImage(w, r, img)
}

func images(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	vars := mux.Vars(r)
	img := "./static/images/" + vars["img"]
	serveImage(w, r, img)
}

func scripts(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	r.Header.Set("Content-Type", "application/javascript")
	vars := mux.Vars(r)
	script := "./static/scripts/" + vars["script"]
	http.ServeFile(w, r, script)
}

func styles(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	r.Header.Set("Content-Type", "text/css")
	vars := mux.Vars(r)
	sheet := "./static/css/" + vars["sheet"]
	http.ServeFile(w, r, sheet)
}

func months(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	months := localization.MonthsByLanguage(config.CurrentLanguage)
	serialized, _ := json.Marshal(months)
	w.WriteHeader(http.StatusOK)
	w.Write(serialized)
}

func reports(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	vars := mux.Vars(r)
	year, err := parseVarYear(vars)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	month, err := parseVarMonth(vars)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	file, err := NewReport(year, month)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	err = export(reportFilename, file, config.CurrentLanguage)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	r.Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	http.ServeFile(w, r, reportFilename)
}
