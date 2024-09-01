package dao

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xeaser/pismo/internal/models"
)

func TestCreateAccount(t *testing.T) {
	dao := NewAccountsDaoer()
	account := &models.Account{Id: 1, DocumentNumber: "12345678"}

	id, err := dao.CreateAccount(account)
	assert.NoError(t, err)

	accountDetails, _ := dao.GetAccountById(id)
	assert.Equal(t, account, accountDetails)
}

func TestGetAccountById(t *testing.T) {
	dao := NewAccountsDaoer()
	account := &models.Account{Id: 1, DocumentNumber: "12345678"}
	id, _ := dao.CreateAccount(account)

	result, err := dao.GetAccountById(id)
	assert.NoError(t, err)
	assert.Equal(t, account, result)

	result, err = dao.GetAccountById(2)
	assert.Nil(t, result)
	assert.NoError(t, err)
}

func TestCreateTransaction(t *testing.T) {
	dao := NewAccountsDaoer()
	transaction := &models.Transaction{
		TransactionId:          1,
		AccountId:              1,
		OperationType_ID:       models.CreditVoucher,
		Amount:                 50,
		EventDateUnixTimestamp: int(time.Now().Unix()),
	}

	id, err := dao.CreateTransaction(transaction)
	assert.NoError(t, err)

	transactionDetails, _ := dao.GetTransactionById(id)
	assert.Equal(t, transaction, transactionDetails)
}

func TestGetTransactionById(t *testing.T) {
	dao := NewAccountsDaoer()
	transaction := &models.Transaction{
		TransactionId:          1,
		AccountId:              1,
		OperationType_ID:       models.CreditVoucher,
		Amount:                 50,
		EventDateUnixTimestamp: int(time.Now().Unix()),
	}
	id, _ := dao.CreateTransaction(transaction)

	result, err := dao.GetTransactionById(id)
	assert.NoError(t, err)
	assert.Equal(t, transaction, result)

	result, err = dao.GetTransactionById(2)
	assert.Nil(t, result)
	assert.NoError(t, err)
}

func TestGetTransactions(t *testing.T) {
	dao := NewAccountsDaoer()
	transaction1 := &models.Transaction{
		TransactionId:          1,
		AccountId:              1,
		OperationType_ID:       models.CreditVoucher,
		Amount:                 50,
		EventDateUnixTimestamp: int(time.Now().Unix()),
	}
	transaction2 := &models.Transaction{
		TransactionId:          1,
		AccountId:              1,
		OperationType_ID:       models.NormalPurchase,
		Amount:                 50,
		EventDateUnixTimestamp: int(time.Now().Unix()),
	}
	_, _ = dao.CreateTransaction(transaction1)
	_, _ = dao.CreateTransaction(transaction2)

	result, err := dao.GetTransactions()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
}
