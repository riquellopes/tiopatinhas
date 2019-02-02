package duck

import (
	"fmt"
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

// Uncle -
type Uncle struct {
	GoalCents   int
	MonthBudget int
}

// Alert -
func (u *Uncle) Alert() (string, error) {
	// // os.Getenv("DUCK_GOAL_CENTS")
	// strconv.Atoi(os.Getenv("DUCK_MONTH_BUDGET_CENTS"))
	// Goal monthly that I need to have every months.
	goalByMonth := u.GoalCents / 12

	// Call service to get all spents of current month.
	spents, err := HowMuchDidISpend()
	moneyLeft := u.MonthBudget + spents

	if spents == 0 || moneyLeft > goalByMonth {
		return "You are walking in the right way.", nil
	}
	// Calc to get monthly variaction.
	variaction := 100 - ((moneyLeft * 100) / goalByMonth)
	// return "You are 10% off your monthly goal!", err
	return fmt.Sprintf("You are %d%% off your monthly goal!", variaction), err
}
