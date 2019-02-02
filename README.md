[![Build Status](https://travis-ci.org/riquellopes/tiopatinhas.svg?branch=master)](https://travis-ci.org/riquellopes/tiopatinhas)
[![Coverage Status](https://coveralls.io/repos/github/riquellopes/tiopatinhas/badge.svg?branch=master)](https://coveralls.io/github/riquellopes/tiopatinhas?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/riquellopes/tiopatinhas)](https://goreportcard.com/report/github.com/riquellopes/tiopatinhas)

Tio Patinhas
============

If you've problems to collect money, I'll give you a help. Because the problem isn't how much do you earn, but how much do you spend.

![Tio patinhas](tio-patinhas.jpg)

#### How can I'll help you?
* **First step**:
It's so simple, but necessary who you'll create an account on [organizze](https://www.organizze.com.br/). Why you need do that? Because the [organizze](https://www.organizze.com.br/) has an [api](https://en.wikipedia.org/wiki/Application_programming_interface) and It'll help you to organize your financial.
* **Second step:** You need to read that simple Portuguese [steps](https://github.com/organizze/api-doc#fazendo-uma-requisi%C3%A7%C3%A3o) to create your access token.
* **Third step:** You should to create an account on [sendgrid](https://sendgrid.com/) service, after account created, access the [documentation](https://sendgrid.com/docs/ui/account-and-settings/api-keys/) to create your app key is very simple.
* **Fourth step:** To scheduled your app, you need to create an [aws account](https://docs.aws.amazon.com/AmazonSimpleDB/latest/DeveloperGuide/AboutAWSAccounts.html), to use the lambda service. To build the lambda deploy, I used the [serverless](https://serverless.com/), because it's simple and I don't want to worry it.

##### Enviroment variables:
```sh
 $ export DUCK_MONTH_BUDGET_CENTS=100
 $ export DUCK_GOAL_CENTS=100
 $ export DUCK_FROM=duck@example.com
 $ export DUCK_TO=duck@example.com
 $ export SENDGRID_KEY=xxxx
 $ export ORGANIZZE_USER=duck@example.com
 $ export ORGANIZZE_PASSWORD_KEY=xxxx
 $ export ORGANIZZE_NAME="First Name Last Name"
```
#### TODO LIST:
* to create another test cases.
* to refactor for create a simple app.
