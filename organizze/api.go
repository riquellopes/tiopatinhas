package organize

import (
	"encoding/json"
	"errors"
	"net/http"
)

// URI -
const URI = "https://api.organizze.com.br/rest/v2"

// Organize -
type Organize struct {
	User        string
	PasswordKey string
}

// ResponseBase -
type ResponseBase interface {
}

// Request -
func Request(service string, object ResponseBase) error {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", URI, nil)
	request.SetBasicAuth("", "")
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("server error")
	}

	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(&object)
}

// GetTransactions -
func (o *Organize) GetTransactions(startDate, endDate string) {

}
