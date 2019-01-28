package organize

import (
	"encoding/json"
	"net/http"
)

// URI -
const URI = "https://api.organizze.com.br/rest/v2"

// Organizze -
type Organizze struct {
	User        string
	PasswordKey string
}

// ResponseBase -
type ResponseBase interface {
}

// Request -
func Request(method, service string, object ResponseBase, organizze Organizze) error {
	client := &http.Client{}
	request, _ := http.NewRequest(method, URI, nil)
	request.SetBasicAuth(organizze.User, organizze.PasswordKey)
	request.Header.Set("User-Agent", "")
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(&object)
}

// GetTransactions -
func (o *Organizze) GetTransactions(startDate, endDate string) {

}
