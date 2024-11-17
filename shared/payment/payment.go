package payment

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	models "github.com/OucheneMohamedNourElIslem658/restaurent/models"
// 	tools "github.com/OucheneMohamedNourElIslem658/restaurent/tools"
// )

var instance Config

func Init() {
	instance = envs
}

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateCustomer(email string, name string) (customer *Customer, err error) {
	client := resty.New()

	requestBody := Customer{
		Email: email,
		Name: name,
	}

	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %v", instance.SecretKey)).
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(instance.BaseURL)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to create customer")
	}

	var createdCustomer *Customer
	if err := json.Unmarshal(resp.Body(), createdCustomer); err != nil {
		return nil, err
	}

	return customer, nil
}