package main

import (
	"database/sql"
	"errors"
)

type product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}