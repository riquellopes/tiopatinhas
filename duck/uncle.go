package duck

import (
	"fmt"
	"os"

	org "github.com/riquellopes/tiopatinhas/organizze"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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
	return fmt.Sprintf("You are %d%% off your monthly goal!", variaction), err
}

// Quack -
func (u *Uncle) Quack() {
	from := mail.NewEmail("Tio Patinhas", os.Getenv("DUCK_FROM"))
	to := mail.NewEmail("Henrique Lopes", os.Getenv("DUCK_TO"))

	goalAlert, _ := u.Alert()
	htmlContent := fmt.Sprintf("<strong>%s</strong>", goalAlert)

	message := mail.NewSingleEmail(from, "Tio Patinhas - Quack Quack", to, goalAlert, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_KEY"))
	_, err := client.Send(message)

	fmt.Println(err)
}
