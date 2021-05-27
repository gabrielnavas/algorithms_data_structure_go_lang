package main

import "fmt"

type Account struct {
	id          string
	accountType string
}

func (account *Account) create(accountType string) *Account {
	account.accountType = accountType
	return account
}

func (account *Account) getById(id string) *Account {
	fmt.Println("getting account by Id")
	return account
}

func (account *Account) deleteById(id string) {
	fmt.Println("delete account by id")
}
