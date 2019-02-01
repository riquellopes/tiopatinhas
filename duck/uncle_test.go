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
