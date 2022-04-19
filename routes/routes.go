package routes

import (
	"github.com/gorilla/mux"
	"github.com/umangraval/Go-Mongodb-REST-boilerplate/controllers"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/trivia", controllers.GetTriviaEndpoint).Methods("GET")
	return router
}
