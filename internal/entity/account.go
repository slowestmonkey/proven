package entity

type AccountID string

type AccountNumber string

type Currency string

type Account struct {
	ID       AccountID
	Number   AccountNumber
	Currency Currency
	Balance  struct{}
}
