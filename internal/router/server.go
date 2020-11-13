package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cobnalt/Go/internal/service"
	"github.com/gorilla/mux"
)

// Server struct
type Server struct {
	router  *mux.Router
	service service.Service
}

// New create new Server instance
func New(service service.Service) *Server {
	return &Server{
		router:  mux.NewRouter(),
		service: service,
	}
}

// Run server
func (s *Server) Run() {
	s.initHandlers()
	err := http.ListenAndServe("localhost:8181", nil)
	if err != nil {
		fmt.Println("Error Server Run")
	}

}

func (s *Server) initHandlers() {
	s.router.HandleFunc("/products/{id:[0-9]+}", s.GetProductByID).Methods("GET")
}

// GetProductByID func
func (s *Server) GetProductByID(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	fmt.Println("id")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Println("There was an error in your request")
		return
	}
	pr, err := s.service.GetProductByID(r.Context(), id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(&pr)
}
