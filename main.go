package main

import (
	"gdbank/cmd/api"
)

type Account struct {
	UUID     string
	Username string
	Password *string
	Balance  float64
}

var accounts = make([]Account, 2)

func main() {
	server := api.NewAPIServer(":8080", nil)

	err := server.Run()
	if err != nil {
		return
	}
}
