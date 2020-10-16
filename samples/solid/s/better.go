package s

import (
	"errors"
	"math/rand"
)

// Single Responsibility Principle -- Improved --

type Iban string

type Account struct {
	accountNumber  Iban
	BalanceInCents int
}

type User struct {
	Accounts []Account
	Password string
}

type AuthenticationService struct { // login and password of every customers are in memory!!!
	UsersByLoginId map[string]User
}

type sessionToken struct {
	Token int
	Owner User
}

func newSessionToken(owner User) *sessionToken {
	return &sessionToken{
		Token: rand.Int(),
		Owner: owner,
	}
}

func (as *AuthenticationService) Authenticate(login string, password string) (*sessionToken, error) {
	if user, exist := as.UsersByLoginId[login]; exist {
		if password == user.Password {
			return newSessionToken(user), nil // If authentication succeeds, return a session token
		}
	}
	return nil, errors.New("Authentication failed") // User does not exist
}

func (a *Account) TransferMoneyTo(token *sessionToken, receivingParty Iban, amount int) error {
	//Check sessionToken
	//Perform the transfer
	return nil
}
