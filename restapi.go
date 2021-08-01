package main

import (
	"fmt"
	"net/http"
)

type customer struct {
	Id     int
	name   string
	gender string
}

func main() {
	http.HandleFunc("/", custinfo)
	http.HandleFunc("/createmap", createmap)
	http.HandleFunc("/update", updatemap)
	http.ListenAndServe(":8080", nil)
}
func createmap(w http.ResponseWriter, r *http.Request) {
	cust := customer{
		Id:     5,
		name:   "John",
		gender: "Male",
	}
	fmt.Fprintf(w, "customer details:", cust)

}
func updatemap(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi John")
}
func custinfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Quest")
}
