package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ApiError struct {
	err    string
	Status int
}

func (e ApiError) Error() string {
	return e.err
}

type APIServer struct {
	store      Storage
	listenAddr string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(ApiError); ok {
				WriteJSON(w, e.Status, e)
				return
			}
			WriteJSON(w, http.StatusBadRequest, ApiError{err: err.Error()})
		}
	}
}

func NewServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {

	router := chi.NewRouter()

	router.Post("/account", makeHTTPHandleFunc(s.handleCreateAccount))

	log.Println("JSON API SERVER RUNNING ON PORT:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, "Online")
}
