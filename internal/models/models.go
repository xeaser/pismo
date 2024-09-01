package models

// OperationType_ID represents the type of operation that can be performed.
type OperationType_ID int

// Constants for defining the types of operations.
const (
	Invalid OperationType_ID = iota - 1
	// All transaction operation types
	All

	// Debit transaction types
	NormalPurchase
	InstallmentPurchase
	Withdrawal

	// Credit Transaction type
	CreditVoucher
)

// GetOperationType returns the OperationType_ID based on the provided integer value.
func GetOperationType(oType int) OperationType_ID {
	switch oType {
	case int(All):
		return All
	case int(NormalPurchase):
		return NormalPurchase
	case int(InstallmentPurchase):
		return InstallmentPurchase
	case int(Withdrawal):
		return Withdrawal
	case int(CreditVoucher):
		return CreditVoucher
	default:
		return Invalid
	}
}

// Account represents a customer account.
type Account struct {
	Id             int    `json:"id"`
	DocumentNumber string `json:"document_number"`
}

// Transaction represents a financial transaction.
type Transaction struct {
	TransactionId          int              `json:"transaction_id"`
	AccountId              int              `json:"account_id"`
	OperationType_ID       OperationType_ID `json:"operationType_id"`
	Amount                 float64          `json:"amount"`
	EventDateUnixTimestamp int              `json:"eventDateTimestamp"`
}

// TransactionFilter represents a filter for transactions.
type TransactionFilter struct {
	TransactionId int `json:"transaction_id"`
	AccountId     int `json:"account_id"`
	OperationType int `json:"operation_type"`
}
