package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/xeaser/pismo/internal/dao"
	"github.com/xeaser/pismo/internal/helper"
	"github.com/xeaser/pismo/internal/models"
)

// AccountHandler defines the interface for handling account operations.
type AccountHandler interface {
	createAccount(w http.ResponseWriter, r *http.Request)
	getAccount(w http.ResponseWriter, r *http.Request)

	createTransaction(w http.ResponseWriter, r *http.Request)
	getTransactionsByFilter(w http.ResponseWriter, r *http.Request)
}

// NewAccountHandler creates a new instance of AccountOperation.
func NewAccountHandler(dao dao.AccountsDaoer) AccountHandler {
	return &AccountOperation{
		dao: dao,
	}
}

// AccountOperation represents the operations that can be performed on accounts.
type AccountOperation struct {
	dao dao.AccountsDaoer
}

// createAccount handles the creation of a new account.
func (ao *AccountOperation) createAccount(w http.ResponseWriter, r *http.Request) {
	if !helper.ValidateHttpMethod(w, r, http.MethodPost) {
		return
	}

	var account models.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		err := fmt.Errorf("invalid request body")
		helper.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	id, err := ao.dao.CreateAccount(&account)
	if err != nil {
		helper.RespondWithStatus(w, http.StatusInternalServerError)
		return
	}

	accountDetails, err := ao.dao.GetAccountById(id)
	if err != nil {
		helper.RespondWithStatus(w, http.StatusInternalServerError)
		return
	}
	helper.RespondWithData(w, accountDetails)
}

// getAccount retrieves an account by its ID.
func (ao *AccountOperation) getAccount(w http.ResponseWriter, r *http.Request) {
	if !helper.ValidateHttpMethod(w, r, http.MethodGet) {
		return
	}

	accountID := r.PathValue("accountId")
	id, err := strconv.Atoi(accountID)
	if err != nil {
		err := fmt.Errorf("invalid account Id")
		helper.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	accountDetails, err := ao.dao.GetAccountById(id)
	if err != nil {
		helper.RespondWithStatus(w, http.StatusInternalServerError)
		return
	}
	helper.RespondWithData(w, accountDetails)
}

// createTransaction handles the creation of a new transaction.
func (ao *AccountOperation) createTransaction(w http.ResponseWriter, r *http.Request) {
	if !helper.ValidateHttpMethod(w, r, http.MethodPost) {
		return
	}

	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		err := fmt.Errorf("invalid request body")
		helper.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	account, err := ao.dao.GetAccountById(transaction.AccountId)
	if err != nil {
		helper.RespondWithStatus(w, http.StatusInternalServerError)
		return
	} else if account == nil {
		err := fmt.Errorf("account doesn't exist")
		helper.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	operationType := models.GetOperationType(int(transaction.OperationType_ID))
	if operationType == models.Invalid || operationType == models.All {
		err := fmt.Errorf("operation type not supported")
		helper.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	transaction.UpdateAmountByOperationType()
	id, err := ao.dao.CreateTransaction(&transaction)
	if err != nil {
		helper.RespondWithStatus(w, http.StatusInternalServerError)
		return
	}

	options := []models.TransactionOption{
		models.WithTransactionId(id),
		models.WithAccountId(transaction.AccountId),
		models.WithOperationType(transaction.OperationType_ID),
	}
	transactionDetails, err := ao.dao.GetTransactions(options...)
	if err != nil {
		helper.RespondWithStatus(w, http.StatusInternalServerError)
		return
	}

	helper.RespondWithData(w, transactionDetails)
}

// getTransactionsByFilter retrieves transactions based on filters.
func (ao *AccountOperation) getTransactionsByFilter(w http.ResponseWriter, r *http.Request) {
	if !helper.ValidateHttpMethod(w, r, http.MethodPost) {
		return
	}

	var filters models.TransactionFilter
	if err := json.NewDecoder(r.Body).Decode(&filters); err != nil {
		err := fmt.Errorf("invalid filters")
		helper.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	options := []models.TransactionOption{}
	if filters.AccountId > 0 {
		options = append(options, models.WithAccountId(filters.AccountId))
	} else {
		err := fmt.Errorf("invalid account id")
		helper.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	operationType := models.GetOperationType(filters.OperationType)
	if operationType != models.Invalid {
		options = append(options, models.WithOperationType(operationType))
	} else {
		err := fmt.Errorf("invalid operation type")
		helper.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	transactions, err := ao.dao.GetTransactions(options...)
	if err != nil {
		helper.RespondWithStatus(w, http.StatusInternalServerError)
		return
	}

	helper.RespondWithData(w, transactions)
}
