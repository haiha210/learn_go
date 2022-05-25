package routes

import (
	"github.com/gorilla/mux"
	"github.com/haiha210/learn_go/11-project-free-code-camp/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandlerFunc("/books/", controllers.CreateBook).Methods("POST")
}
