package main

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
	

	"github.com/gorilla/mux"
)
type API struct {
    listenAddr string
	store Storage
}

func NewAPI(listenAddr string, store Storage) *API {
    return &API{
        listenAddr: listenAddr,
		store: store,
    }
}

func (s *API) Run() {
    router := mux.NewRouter()
    router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount)).Methods("GET", "POST", "DELETE")
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(s.handleGetAccountByID)).Methods("GET")
    
	log.Println("Listening on", s.listenAddr)
    http.ListenAndServe(s.listenAddr, router)
}

func (s *API) handleAccount(w http.ResponseWriter, r *http.Request) error{
	// handle account
	if r.Method == "GET"{
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST"{
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE"{
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed")
}
func (s *API) handleGetAccount(w http.ResponseWriter, r *http.Request) error{
	// handle account
	accounts, err := s.store.GetAccounts()
	if err != nil{
		return err
	}
	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *API) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error{
	// handle account
	id := mux.Vars(r)["id"]
	fmt.Println("ID:", id)
	return WriteJSON(w, http.StatusOK, &Account{})
}
func (s *API) handleCreateAccount(w http.ResponseWriter, r *http.Request) error{
	// handle account
	CreateAccountReq := new(CreateAccountRequest)
	// CreateAccountReq := CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(CreateAccountReq); err != nil{
		return err
	}

	account := NewAccount(CreateAccountReq.FirstName, CreateAccountReq.LastName, CreateAccountReq.Email, CreateAccountReq.Number, CreateAccountReq.Balance)
	if err := s.store.CreateAccount(account); err != nil{
		return err
	}
	return WriteJSON(w, http.StatusOK, account)
}
func (s *API) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error{
	// handle account
	return nil
}
func (s *API) handleTransferAccount(w http.ResponseWriter, r *http.Request) error{
	// handle account
	return nil
}

// helper function

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
    return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error
type ApiError struct {
	Error string
}
func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if err := f(w, r); err != nil{
			// handler error
			WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
			
		}
	}
}

