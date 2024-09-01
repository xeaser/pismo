package account

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mockDao "github.com/xeaser/pismo/internal/dao/mocks"
	"github.com/xeaser/pismo/internal/models"
)

// TestCreateAccount tests the createAccount method.
func TestCreateAccount(t *testing.T) {
	mockDao := new(mockDao.AccountsDaoer)
	handler := NewAccountHandler(mockDao)

	account := models.Account{
		Id:             1,
		DocumentNumber: "12345678",
	}
	body, _ := json.Marshal(account)

	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	mockDao.On("CreateAccount", &account).Return(1, nil)
	mockDao.On("GetAccountById", 1).Return(&account, nil)

	handler.createAccount(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Account
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, account, response)
}

// TestGetAccount tests the getAccount method.
func TestGetAccount(t *testing.T) {
	mockDao := new(mockDao.AccountsDaoer)
	handler := NewAccountHandler(mockDao)

	account := models.Account{
		Id:             1,
		DocumentNumber: "12345678",
	}
	mockDao.On("GetAccountById", 1).Return(&account, nil)

	req := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
	// currently this test fails as http Serve mux can't be mocked easily
	// can we solved if we use query params adn set it in test
	w := httptest.NewRecorder()

	handler.getAccount(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Account
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, account, response)
}

// TestCreateTransaction tests the createTransaction method.
func TestCreateTransaction(t *testing.T) {
	mockDao := new(mockDao.AccountsDaoer)
	handler := NewAccountHandler(mockDao)

	transactionTime := int(time.Now().Unix())
	transaction := models.Transaction{
		TransactionId:          1,
		AccountId:              1,
		OperationType_ID:       1,
		Amount:                 20.5,
		EventDateUnixTimestamp: transactionTime,
	}
	body, _ := json.Marshal(transaction)

	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	mockDao.On("GetAccountById", 1).Return(&models.Account{
		Id:             1,
		DocumentNumber: "12345678",
	}, nil)

	mockDao.On("CreateTransaction", &transaction).Return(1, nil)
	mockDao.On("GetTransactions", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Transaction{&transaction}, nil)
	transaction.UpdateAmountByOperationType()

	handler.createTransaction(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedTransaction := []*models.Transaction{
		{
			TransactionId:          1,
			AccountId:              1,
			OperationType_ID:       1,
			Amount:                 -20.5,
			EventDateUnixTimestamp: transactionTime,
		},
	}
	var response []*models.Transaction
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, expectedTransaction, response)
}

// TestGetTransactionsByFilter tests the getTransactionsByFilter method.
func TestGetTransactionsByFilter(t *testing.T) {
	mockDao := new(mockDao.AccountsDaoer)
	handler := NewAccountHandler(mockDao)

	filters := models.TransactionFilter{AccountId: 1}
	body, _ := json.Marshal(filters)

	req := httptest.NewRequest(http.MethodPost, "/transactions/filter", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	// can be updated as per requirement
	mockDao.On("GetTransactions", mock.Anything, mock.Anything).Return([]*models.Transaction{}, nil)

	handler.getTransactionsByFilter(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
