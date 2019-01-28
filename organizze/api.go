package organizze

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// URI -
const URI = "https://api.organizze.com.br/rest/v2"

// Organizze -
type Organizze struct {
	User        string
	PasswordKey string
	Name        string
}

// ResponseBase -
type ResponseBase interface {
}

// Request -
func Request(method, service string, object ResponseBase, organizze *Organizze) error {
	client := &http.Client{}
	request, _ := http.NewRequest(method, service, nil)
	request.SetBasicAuth(organizze.User, organizze.PasswordKey)
	request.Header.Set("User-Agent", fmt.Sprintf("%s (%s)", organizze.Name, organizze.User))
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(&object)
}

// GetTransactions -
func (o *Organizze) GetTransactions(date ...string) ([]Transaction, error) {
	transaction := []Transaction{}
	service := fmt.Sprintf("%s/%s", URI, "transactions/")

	err := Request("GET", service, &transaction, o)

	return transaction, err
}
