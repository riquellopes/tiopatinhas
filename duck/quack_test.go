package duck

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func Test_should_send_a_quack_to_email_defined(t *testing.T) {
	uncle := Uncle{
		// R$ 20.000,00
		GoalCents: 2000000,
		// R$ 5.000,00
		MonthBudget: 500000,
	}

	gock.New("https://api.sendgrid.com").
		Post("/v3/mail/send").
		Reply(200)

	uncle.Quack()
	assert.True(t, gock.IsDone())
}
