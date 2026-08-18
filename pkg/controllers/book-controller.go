package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jeetdas5/go-bookstore/pkg/models"
	"github.com/jeetdas5/go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks, err := models.GetAllBooks()
	if err != nil {
		fmt.Println(err)
		return
	}

	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error parsing")
		return
	}
	book, _ := models.GetBookById(id)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()

	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error parsing")
		return
	}
	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
		return
	}
	book, db := models.GetBookById(id)

	if db == nil {
		fmt.Println("No book found")
		return
	}

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}

	if updateBook.Title != "" {
		book.Title = updateBook.Title
	}

	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}


	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	db.Save(&book)

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
