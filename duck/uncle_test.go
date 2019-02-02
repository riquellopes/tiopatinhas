package duck

import (
	"testing"

	org "github.com/riquellopes/tiopatinhas/organizze"
	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func Test_should_get_minus_178293_when_to_call_the_func(t *testing.T) {
	defer gock.Off()

	response := []org.Transaction{
		{ID: 95, Description: "Just a little cafe on mall.", AmountCents: -22999},
		{ID: 96, Description: "Beer after the work.", AmountCents: -155294},
	}

	gock.New(org.URI).
		Get("/transactions").
		Reply(200).
		JSON(response)

	money, _ := HowMuchDidISpend()

	assert.Equal(t, money, -178293)
}

func Test_should_get_0_when_call_the_func(t *testing.T) {
	defer gock.Off()

	response := []org.Transaction{}

	gock.New(org.URI).
		Get("/transactions").
		Reply(200).
		JSON(response)

	money, _ := HowMuchDidISpend()

	assert.Equal(t, money, 0)
}

func Test_should_get_10_porcent_msg(t *testing.T) {
	defer gock.Off()

	response := []org.Transaction{
		// R$ 500,00
		{ID: 1, AmountCents: -50000},
		// R$ 500,00
		{ID: 2, AmountCents: -50000},
		// R$ 1000,00
		{ID: 3, AmountCents: -100000},
		// R$ 2000,00
		{ID: 4, AmountCents: -200000},
	}

	gock.New(org.URI).
		Get("/transactions").
		Reply(200).
		JSON(response)

	uncle := Uncle{
		// R$ 20.000,00
		GoalCents: 2000000,
		// R$ 5.000,00
		MonthBudget: 500000,
	}

	way, _ := uncle.Alert()
	assert.Equal(t, way, "You are 40% off your monthly goal!")
}

func Test_should_get_ok_msg(t *testing.T) {
	defer gock.Off()

	response := []org.Transaction{}

	gock.New(org.URI).
		Get("/transactions").
		Reply(200).
		JSON(response)

	uncle := Uncle{
		// R$ 20.000,00
		GoalCents: 2000000,
		// R$ 5.000,00
		MonthBudget: 500000,
	}

	way, _ := uncle.Alert()

	assert.Equal(t, way, "You are walking in the right way.")
}

func Test_should_get_ok_msg_when_financial_ok(t *testing.T) {
	defer gock.Off()

	response := []org.Transaction{
		// R$ 200,00
		{ID: 95, Description: "Just a little cafe on mall.", AmountCents: -20000},
	}

	gock.New(org.URI).
		Get("/transactions").
		Reply(200).
		JSON(response)

	uncle := Uncle{
		// R$ 20.000,00
		GoalCents: 2000000,
		// R$ 5.000,00
		MonthBudget: 500000,
	}

	way, _ := uncle.Alert()

	assert.Equal(t, way, "You are walking in the right way.")
}
