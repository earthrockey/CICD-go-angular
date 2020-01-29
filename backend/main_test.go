package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/earthrockey/CICD-go-angular/backend/model"
)

func TestOpenServer(t *testing.T) {
	log.Print("The service is start on http://localhost" + getPort())
	go HandleRequest()
}

func TestGetAllBook(t *testing.T) {
	log.Println("Test Get All Book")
	r, err := http.Get("http://localhost:8080/api/get/allbook")
	if err != nil {
		log.Print(err)
	}
	defer r.Body.Close()
	var books []model.Book
	err = json.NewDecoder(r.Body).Decode(&books)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(books)
}

func TestCreateBook(t *testing.T) {
	log.Println("Test Create Book")
	// r, err := http.Get("http://localhost:8080/api/get/allbook")
	// if err != nil {
	// 	log.Print(err)
	// }
	// defer r.Body.Close()
	// var books []model.Book
	// err = json.NewDecoder(r.Body).Decode(&books)
	// if err != nil {
	// 	log.Print(err)
	// }
	// fmt.Println(books)  
}
