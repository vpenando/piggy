package routing

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/vpenando/piggy/piggy"
)

func getCategories(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	categories, err := categoryController.ReadAll()
	handleError(w, err, http.StatusInternalServerError)
	serialized, _ := json.Marshal(categories)
	r.Header.Set("Content-Type", "application/json")
	w.Write(serialized)
}

// Possible HTTP errors:
//  * Invalid request body => StatusUnprocessableEntity
//  * Invalid icon file (PNG expected) => StatusUnprocessableEntity
//  * Icon saving failure => StatusInternalServerError
//  * Category saving failure => StatusInternalServerError
func postCategory(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	rawCategory := struct {
		Name string
		Icon []byte
	}{}
	err = json.Unmarshal(body, &rawCategory)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	if !isPNG(rawCategory.Icon) {
		handleError(w, errors.New("not a PNG file"), http.StatusUnprocessableEntity)
		return
	}
	filename := "./categories/custom_" + strings.ReplaceAll(rawCategory.Name, " ", "_") + ".png"
	err = saveCategoryIcon(rawCategory.Icon, filename)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	category := piggy.NewCategory(rawCategory.Name, filename)
	_, err = categoryController.Create(category)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	r.Header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
