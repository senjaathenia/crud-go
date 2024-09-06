package main

import (
	"go-native/config"
	"go-native/controllers/categorycontroller"
	"go-native/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	//HomePage
	http.HandleFunc("/", homecontroller.Welcome)

	//Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)
	
	config.ConnectDB()
	log.Println("Server is running on port 8080")
	http.ListenAndServe(": 8080", nil)
}