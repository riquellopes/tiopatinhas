package organizze

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func Test_should_get_a_list_of_itens(t *testing.T) {
	defer gock.Off()

	response := []Transaction{
		{ID: 95, Description: "Just a little cafe on mall."},
		{ID: 96, Description: "Beer after the work."},
	}

	gock.New(URI).
		Get("/transactions").
		Reply(200).
		JSON(response)

	oozz := Organizze{User: "lopes@example.com", PasswordKey: "xxxxxxxxxxxxxxxx", Name: "Rick"}
	result, _ := oozz.GetTransactions()

	assert.Equal(t, result, response)
}
