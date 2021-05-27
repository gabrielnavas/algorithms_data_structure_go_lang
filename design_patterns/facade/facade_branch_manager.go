package main

type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{
		&Account{},
		&Customer{},
		&Transaction{},
	}
}

func (facade *BranchManagerFacade) CreateCustomerAccount(
	customerName string,
	accountType string) (*Customer, *Account) {
	customer := facade.customer.create(customerName)
	account := facade.account.create(accountType)
	return customer, account
}

func (facade *BranchManagerFacade) CreateTransaction(srcAccountId string, destAccountId string, amount float32) *Transaction {
	transaction := facade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction
}
