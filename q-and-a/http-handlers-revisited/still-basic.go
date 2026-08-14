package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// User represents a person in our system.
type User struct {
	Name string `json:"name"`
}

// UserService provides ways of working with users.
type UserService interface {
	Register(user User) (insertedID string, err error)
}

// UserServer provides an HTTP API for working with users.
type UserServer struct {
	service UserService
}

// NewUserServer creates a UserServer.
func NewUserServer(service UserService) *UserServer {
	return &UserServer{service}
}

// RegisterUser is a http handler for storing users.
func (u *UserServer) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	id, err := u.service.Register(user)
	if err != nil {
		http.Error(w, "cannot register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, id)
}

type InMemoryUserService struct {
	users []User
}

func (i *InMemoryUserService) Register(user User) (insertedID string, err error) {
	i.users = append(i.users, user)
	return strconv.Itoa(len(i.users) - 1), nil
}

func main() {
	srv := NewUserServer(&InMemoryUserService{})
	log.Println("starting server...")
	log.Println(http.ListenAndServe(":8080", http.HandlerFunc(srv.RegisterUser)))
}
