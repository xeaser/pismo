package account

import (
	"net/http"

	"github.com/xeaser/pismo/internal/dao"
)

// RegisterHandler registers the account handlers with the given ServeMux.
func RegisterHandler(mux *http.ServeMux) *http.ServeMux {
	ah := NewAccountHandler(dao.NewAccountsDaoer())
	mux.HandleFunc("/accounts", ah.createAccount)
	mux.HandleFunc("/accounts/{accountId}", ah.getAccount)
	mux.HandleFunc("/transactions", ah.createTransaction)
	mux.HandleFunc("/transactionsByFilter", ah.getTransactionsByFilter)
	return mux
}
