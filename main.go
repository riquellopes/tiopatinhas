package main

import (
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	un "github.com/riquellopes/tiopatinhas/duck"
)

func quack() (string, error) {
	var (
		GoalCents, _   = strconv.Atoi(os.Getenv("DUCK_GOAL_CENTS"))
		MonthBudget, _ = strconv.Atoi(os.Getenv("DUCK_MONTH_BUDGET_CENTS"))
	)

	uncle := un.Uncle{
		GoalCents:   GoalCents,
		MonthBudget: MonthBudget,
	}

	uncle.Quack()

	return "Okay", nil
}

func main() {
	lambda.Start(quack)
}
