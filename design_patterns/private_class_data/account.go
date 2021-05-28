package main

type AccountDetails struct {
	id, accountType string
}

type Account struct {
	details      *AccountDetails // first letter is lowercase, is private
	CustomerName string
}

func (account *Account) setDetails(id, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

func (account *Account) getId() string {
	return account.details.id
}

func (account *Account) getAccountType() string {
	return account.details.accountType
}
