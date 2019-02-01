package duck

import (
	"os"

	org "github.com/riquellopes/tiopatinhas/organizze"
)

// Organizze -
var Organizze = &org.Organizze{
	User:        os.Getenv("ORGANIZZE_USER"),
	PasswordKey: os.Getenv("ORGANIZZE_PASSWORD_KEY"),
	Name:        os.Getenv("ORGANIZZE_NAME"),
}

// HowMuchDidISpend -
func HowMuchDidISpend() (int, error) {
	transactions, err := Organizze.GetTransactions()
	total := 0
	for _, element := range transactions {
		total += element.AmountCents
	}
	return total, err
}
