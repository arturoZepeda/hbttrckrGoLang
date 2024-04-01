package main

import (
	"fmt"
	"log"
	"net/http"
)

func paginaInicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the homePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users Page")
	fmt.Println("Endpoint Hit: users")
}

func handleRequest() {
	http.HandleFunc("/", paginaInicial)
	http.HandleFunc("/users", getUsers)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequest()
}
