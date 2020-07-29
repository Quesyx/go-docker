package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "HELLO WEB")
    })
	fmt.Println("server is listening")
    http.ListenAndServe(":8080", r)
}
