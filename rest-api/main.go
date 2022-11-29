package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

type Customers []Customer

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Admin. Welcome to System!")
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Admin. Welcome to Customer System!\n")
	c1 := Customers{Customer{Name: "Mehmet", Surname: "SOLAK", Age: 19}}
	json.NewEncoder(w).Encode(c1)
}

func SeeCustomer(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8080/addCustomer")
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	json_body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	customer := Customers{Customer{}}
	json.Unmarshal(json_body, &customer)
	json.NewEncoder(w).Encode(customer)
}

func main() {
	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/addCustomer", AddCustomer)
	http.HandleFunc("/seeCustomer", SeeCustomer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
