package http

import (
	"fmt"
	"net/http"
	"strconv"
	"unbeatable-abayomi/commentsApi/internal/comment"

	"github.com/gorilla/mux"
)

//Handler stores the pointer to our comment service
type Handler struct{
    Router *mux.Router
	Service *comment.Service
}


//NewHandler - returns a pointer to a handler
func NewHandler(service *comment.Service) *Handler{
	return &Handler{
		Service: service,
	}
}


//SetupRoutes -sets up all the routes for our application
func (h *Handler) SetupRoutes(){
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/comment/{id}")
	h.Router.HandleFunc("/api/health", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "I am still alive");
	})
}
//GetComment - will reterive a comment by Id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
    
	i,err := strconv.ParseUint(id, 10,64)

	if err != nil {
		fmt.Fprintf(w,"Unable to Pares UNIT from ID");
	}
	comment, err := h.Service.GetComment(uint(i));
	if err != nil{
		fmt.Fprintf(w, "Error Reteriving comment by Id")
	}
	fmt.Fprintf(w,"%+v", comment)
}