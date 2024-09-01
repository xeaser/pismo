package dao

import (
	models "github.com/xeaser/pismo/internal/models"
)

// AccountsDaoer is an interface that defines the operations that can be performed on accounts and transactions.
type AccountsDaoer interface {
	GetAccountById(id int) (*models.Account, error)
	GetAccounts() ([]*models.Account, error)
	CreateAccount(account *models.Account) (int, error)

	GetTransactionById(id int) (*models.Transaction, error)
	GetTransactions(filters ...models.TransactionOption) ([]*models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) (int, error)
}

// AccountsDao is a struct that implements the AccountsDaoer interface.
type AccountsDao struct {
	Accounts     []*models.Account
	Transactions []*models.Transaction
}

// NewAccountsDaoer creates a new instance of AccountsDaoer.
func NewAccountsDaoer() AccountsDaoer {
	return &AccountsDao{}
}

// GetAccountById retrieves an account by its ID.
func (dao *AccountsDao) GetAccountById(id int) (*models.Account, error) {
	for _, acc := range dao.Accounts {
		if acc.Id == id {
			return acc, nil
		}
	}
	return nil, nil
}

// GetAccounts retrieves all accounts.
func (dao *AccountsDao) GetAccounts() ([]*models.Account, error) {
	return dao.Accounts, nil
}

// CreateAccount creates a new account.
func (dao *AccountsDao) CreateAccount(account *models.Account) (int, error) {
	if len(dao.Accounts) > 0 {
		account.Id = dao.Accounts[len(dao.Accounts)-1].Id + 1
	} else {
		account.Id = 1
	}

	dao.Accounts = append(dao.Accounts, account)
	return account.Id, nil
}

// GetTransactionById retrieves a transaction by its ID.
func (dao *AccountsDao) GetTransactionById(id int) (*models.Transaction, error) {
	for _, t := range dao.Transactions {
		if t.TransactionId == id {
			return t, nil
		}
	}
	return nil, nil
}

// GetTransactions retrieves all transactions based on the provided filters.
func (dao *AccountsDao) GetTransactions(options ...models.TransactionOption) ([]*models.Transaction, error) {
	var result []*models.Transaction
	for _, t := range dao.Transactions {
		match := true
		for _, op := range options {
			if !op(t) {
				match = false
				break
			}
		}
		if match {
			result = append(result, t)
		}
	}
	return result, nil
}

// CreateTransaction creates a new transaction per account
func (dao *AccountsDao) CreateTransaction(transaction *models.Transaction) (int, error) {
	accountTransactions := make([]*models.Transaction, 0)
	for _, t := range dao.Transactions {
		if t.AccountId == transaction.AccountId {
			accountTransactions = append(accountTransactions, t)
		}
	}
	if len(accountTransactions) > 0 {
		transaction.TransactionId = accountTransactions[len(accountTransactions)-1].TransactionId + 1
	} else {
		transaction.TransactionId = 1
	}
	dao.Transactions = append(dao.Transactions, transaction)
	return transaction.TransactionId, nil
}
