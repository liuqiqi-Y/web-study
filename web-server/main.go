package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "----- hello ")
}

func login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintln("have no name"))
	}
	fmt.Fprintf(w, fmt.Sprintf("well come %s", r.Form["name"][0]))
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", login)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
