package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./quickstart-1611068463766-firebase-adminsdk-114sr-3fc0c71eb8.json")
	router := mux.NewRouter()
	const port string = ":8080"
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Up...Running")
		fmt.Fprintln(w, "Hellooooo")
	})
	router.HandleFunc("/get", GetPosts).Methods("GET")
	router.HandleFunc("/add", AddPosts)
	log.Println("Server Listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
