package organizze

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func Test_should_get_a_list_of_itens(t *testing.T) {
	defer gock.Off()

	response := []Transaction{
		{ID: 95, Description: "Just a little cafe on mall.", AmountCents: -22999},
		{ID: 96, Description: "Beer after the work.", AmountCents: -155294},
	}

	gock.New(URI).
		Get("/transactions").
		Reply(200).
		JSON(response)

	oozz := Organizze{User: "lopes@example.com", PasswordKey: "xxxxxxxxxxxxxxxx", Name: "Rick"}
	result, _ := oozz.GetTransactions()

	assert.Equal(t, result, response)
}

func Test_should_get_error(t *testing.T) {
	defer gock.Off()

	gock.New(URI).
		Get("/transactions").
		Reply(200).
		SetError(errors.New("server error"))

	oozz := Organizze{User: "lopes@example.com", PasswordKey: "xxxxxxxxxxxxxxxx", Name: "Rick"}
	_, error := oozz.GetTransactions()

	assert.Error(t, error)
}
