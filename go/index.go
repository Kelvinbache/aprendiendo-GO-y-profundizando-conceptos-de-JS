package main

import (
	"fmt"
	"log"
	"net/http"
)

// abriendo archivos un archivo json

func archivo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

func json(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2000)

	file := r.MultipartForm.File["file"]

	fmt.Println(file)
}

func puerto() {
	http.HandleFunc("/", archivo)
	http.HandleFunc("/file", json)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	puerto()
}
