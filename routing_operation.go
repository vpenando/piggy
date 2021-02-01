package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vpenando/piggy/piggy"
)

// Possible HTTP errors:
//  * URI params not provided => StatusBadRequest
//  * Operations reading failure => StatusInternalServerError
//  * JSON encoding failure => StatusInternalServerError
func getOperations(w http.ResponseWriter, r *http.Request) {
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
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, time.Local).UTC()
	endDate := time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local).UTC()
	operations, err := operationController.ReadAllBetween(startDate, endDate)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	serialized, err := json.Marshal(operations)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	r.Header.Set("Content-Type", "application/json")
	w.Write(serialized)
}

// Possible HTTP errors:
//  * Invalid request body => StatusUnprocessableEntity
//  * Operations saving failure => StatusInternalServerError
func postOperation(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	var operation piggy.Operation
	err = json.Unmarshal(body, &operation)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	log.Println("Creating 1 operation...")
	createdOperation, err := operationController.CreateOne(operation)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	serialized, _ := json.Marshal(createdOperation)
	r.Header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(serialized)
}

// Possible HTTP errors:
//  * Invalid request body => StatusUnprocessableEntity
//  * Operations saving failure => StatusInternalServerError
func postOperations(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	var operations []piggy.Operation
	err = json.Unmarshal(body, &operations)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	log.Println("Creating", len(operations), "operation(s)...")
	createdOperations, err := operationController.CreateMany(operations)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	serialized, _ := json.Marshal(createdOperations)
	r.Header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(serialized)
}

// Possible HTTP errors:
//  * Invalid request body => StatusUnprocessableEntity
//  * Operations saving failure => StatusInternalServerError
func updateOperations(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	var operations []piggy.Operation
	err = json.Unmarshal(bytes, &operations)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	updatedOperations, err := operationController.UpdateMany(operations)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	serialized, _ := json.Marshal(updatedOperations)
	w.WriteHeader(http.StatusOK)
	w.Write(serialized)
}

// Possible HTTP errors:
//  * Invalid request body => StatusUnprocessableEntity
//  * Operations deletion failure => StatusInternalServerError
func deleteOperations(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	var ids []int
	err = json.Unmarshal(body, &ids)
	if err != nil {
		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}
	if err := operationController.DeleteMany(ids); err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	r.Header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
