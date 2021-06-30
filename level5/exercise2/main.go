package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID    string  `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductInventory struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

var Inventory []ProductInventory

func add(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var productInventory ProductInventory

	json.Unmarshal(reqBody, &productInventory)

	Inventory = append(Inventory, productInventory)

	json.NewEncoder(w).Encode(productInventory)

}

func update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	flag := false
	for index, i := range Inventory {
		if i.Product.ID == id {
			Inventory = append(Inventory[:index], Inventory[index+1:]...)
			flag = true
		}
	}
	if !flag {
		w.WriteHeader(http.StatusNotFound)
	} else {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var productInventory ProductInventory

		json.Unmarshal(reqBody, &productInventory)
		productInventory.Product.ID = id

		Inventory = append(Inventory, productInventory)

		json.NewEncoder(w).Encode(productInventory)
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	flag := false
	for index, i := range Inventory {
		if i.Product.ID == id {
			Inventory = append(Inventory[:index], Inventory[index+1:]...)
			flag = true
		}
	}
	if !flag {
		w.WriteHeader(http.StatusNotFound)
	}
}

func getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	flag := false

	for _, i := range Inventory {
		if i.Product.ID == id {
			json.NewEncoder(w).Encode(i)
			flag = true
		}
	}
	if !flag {
		w.WriteHeader(http.StatusNotFound)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	for _, i := range Inventory {
		json.NewEncoder(w).Encode(i)
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", add).Methods("POST")
	myRouter.HandleFunc("/products/{id}", delete).Methods("DELETE")
	myRouter.HandleFunc("/products/{id}", update).Methods("PUT")
	myRouter.HandleFunc("/products", get)
	myRouter.HandleFunc("/products/{id}", getById)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
