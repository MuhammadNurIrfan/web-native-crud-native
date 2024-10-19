package main

import (
	"go-web-native-crud/config"
	"go-web-native-crud/controllers/categorycontroller"
	"go-web-native-crud/controllers/homecontroller"
	"go-web-native-crud/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//kita panggil 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//kita panggil 2. Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// kita panggil 3. Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Serve running on port 8080")
	http.ListenAndServe(":8080", nil)
}
