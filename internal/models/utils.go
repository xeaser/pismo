package models

func (t *Transaction) UpdateAmountByOperationType() {
	switch t.OperationType_ID {
	case NormalPurchase, InstallmentPurchase, Withdrawal:
		t.Amount = -t.Amount
	default:

	}
}
