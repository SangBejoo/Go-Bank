package main

import (
	"database/sql"
	"time"
	"fmt"

	_ "github.com/lib/pq"
)



type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
}


type PostgresStore struct {
	// postgres connection
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStore, error) {
	connStr := "user=root dbname=subagiya1 password=secret host=localhost port=5431 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if (err != nil) {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

// Make sure Init is called during NewPostgresStorage
func (s *PostgresStore) Init() error {
    if err := s.CreateAccountTable(); err != nil {
        return fmt.Errorf("error creating account table: %v", err)
    }
    return nil
}
// Update CreateAccountTable function
func (s *PostgresStore) CreateAccountTable() error {
    query := `
        DROP TABLE IF EXISTS account;
        CREATE TABLE account (
            id SERIAL PRIMARY KEY,
            first_name VARCHAR(50),
            last_name VARCHAR(50),
            email VARCHAR(50),
            number BIGINT,
            balance FLOAT,
            created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
        )`
    _, err := s.db.Exec(query)
    if err != nil {
        return err
    }
    return nil
}

// Update CreateAccount function to match table schema
func (s *PostgresStore) CreateAccount(acc *Account) error {
    query := `
        INSERT INTO account 
        (first_name, last_name, email, number, balance, created_at) 
        VALUES ($1, $2, $3, $4, $5, $6) 
        RETURNING id, first_name, last_name, email, number, balance, created_at`
        
    return s.db.QueryRow(
        query,
        acc.FirstName,
        acc.LastName, 
        acc.Email,
        acc.Number,
        acc.Balance,
        time.Now(),
    ).Scan(
        &acc.ID,
        &acc.FirstName,
        &acc.LastName,
        &acc.Email,
        &acc.Number,
        &acc.Balance,
        &acc.CreatedAt,
    )
}

func (s *PostgresStore) DeleteAccount(id int) error {
	
	return nil
}
func (s *PostgresStore) UpdateAccount(a *Account) error {
	
	return nil
}
func (p *PostgresStore) GetAccountByID(id int) (*Account, error) {
    return nil, nil
}
func (p *PostgresStore) GetAccounts() ([]*Account, error) {
    rows, err := p.db.Query("SELECT * FROM account")
    if err != nil {
        return nil, err
    }
    accounts := []*Account{}
    for rows.Next() {
        account := &Account{}
        if err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Email, &account.Number, &account.Balance, &account.CreatedAt); err != nil {
            return nil, err
        }
        accounts = append(accounts, account)
    }
    return accounts, nil
}
