package routes

import (
	"github.com/abdolrhman/simple-go-api/controllers"
	"github.com/gorilla/mux"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/trivia", controllers.GetTriviaEndpoint).Methods("GET")
	return router
}
