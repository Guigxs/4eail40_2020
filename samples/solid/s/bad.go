package s

// Single Responsibility Principle -- BAD--
type SecuredBankAccount struct {
	Iban     string
	Balance  float64 //€
	Password string
}

func (b *SecuredBankAccount) TransferMoneyTo(password string, receivingParty SecuredBankAccount, amount int) {
	//...
}
