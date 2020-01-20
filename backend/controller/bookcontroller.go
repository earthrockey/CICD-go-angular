package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/earthrockey/CICD-go-angular/backend/model"
)

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var books []model.Book
	file, _ := ioutil.ReadFile("./database/books.json")
	json.Unmarshal(file, &books)
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Print(err)
	}
	var books []model.Book
	file, _ := ioutil.ReadFile("./database/books.json")
	json.Unmarshal(file, &books)
	books = append(books, book)
	jsonString, _ := json.Marshal(books)
	ioutil.WriteFile("./database/books.json", jsonString, os.ModePerm)
	json.NewEncoder(w).Encode(books)
}

func DeleteBookByID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Print(err)
	}
	var books []model.Book
	file, _ := ioutil.ReadFile("./database/books.json")
	json.Unmarshal(file, &books)

	for i, item := range books {
		if item.ID == book.ID {
			copy(books[i:], books[i+1:])
			books = books[:len(books)-1]
			break
		}
	}
	jsonString, _ := json.Marshal(books)
	ioutil.WriteFile("./database/books.json", jsonString, os.ModePerm)
	json.NewEncoder(w).Encode(books)
}

func EditBookByID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Print(err)
	}
	var books []model.Book
	file, _ := ioutil.ReadFile("./database/books.json")
	json.Unmarshal(file, &books)

	for i, item := range books {
		if item.ID == book.ID {
			books[i].Name = book.Name
			books[i].Code = book.Code
			books[i].Detail = book.Detail
			break
		}
	}
	jsonString, _ := json.Marshal(books)
	ioutil.WriteFile("./database/books.json", jsonString, os.ModePerm)
	json.NewEncoder(w).Encode(books)
}
