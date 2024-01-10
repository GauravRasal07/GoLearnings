package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GauravRasal07/go_bookstore/pkg/config"
	"github.com/GauravRasal07/go_bookstore/pkg/models"
	"github.com/GauravRasal07/go_bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var NewBook models.Book

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetBooks(res http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	response, _ := json.Marshal(newBooks)

	res.Header().Set("content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing!")
	}

	bookDetails, _ := models.GetBookById(ID)

	response, _ := json.Marshal(bookDetails)

	res.Header().Set("content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(req, CreateBook)

	book := CreateBook.CreateBook()

	response, _ := json.Marshal(book)

	res.Header().Set("content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(req, updateBook)

	bookId := mux.Vars(req)["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing!")
	}

	bookDetails, _ := models.GetBookById(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	response, _ := json.Marshal(bookDetails)

	res.Header().Set("content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	bookId := mux.Vars(req)["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing!")
	}

	deletedBook := models.DeleteBook(ID)

	response, _ := json.Marshal(deletedBook)

	res.Header().Set("content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}
