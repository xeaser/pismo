package models

// TransactionFilter defines a type for filtering transactions
type TransactionOption func(*Transaction) bool

func WithTransactionId(id int) TransactionOption {
	return func(t *Transaction) bool {
		return t.TransactionId == id
	}
}

// Filters transactions by AccountId
func WithAccountId(accountId int) TransactionOption {
	return func(t *Transaction) bool {
		return t.AccountId == accountId
	}
}

// Filters transactions by OperationType_ID
func WithOperationType(operationType OperationType_ID) TransactionOption {
	return func(t *Transaction) bool {
		if operationType == All {
			return true
		}
		return t.OperationType_ID == operationType
	}
}
