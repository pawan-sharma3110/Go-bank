package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, v interface{}) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
    router.HandleFunc("/account",makeHTTPHandlerFunc(s.handelAccount))
	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handelAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handelGetAccount(w, r)
	case "POST":
		return s.handelCreateAccount(w, r)
	case "DELETE":
		return s.handelDeleteAccount(w, r)
	default:
		return fmt.Errorf("%s Method not allowed", r.Method)
	}
}

func (s *APIServer) handelGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handelCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handelDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handelTransfert(w http.ResponseWriter, r *http.Request) error {
	return nil
}
