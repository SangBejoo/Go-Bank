package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Number    int64   `json:"number"`
	Balance   float64 `json:"balance"`
}

type Account struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Number    int64   `json:"number"`
	Balance   float64 `json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName, email string, number int64, balance float64) *Account {
	return &Account{
		
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Number:    int64(rand.Intn(1000000000000)),
		Balance:   balance,
		CreatedAt:  time.Now(),
	}
}
