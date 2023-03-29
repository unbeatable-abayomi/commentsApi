package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Handler stores the pointer to our comment service
type Handler struct{
    Router *mux.Router
}


//NewHandler - returns a pointer to a handler
func NewHandler() *Handler{
	return &Handler{}
}


//SetupRoutes -sets up all the routes for our application
func (h *Handler) SetupRoutes(){
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "I am still alive");
	})
}