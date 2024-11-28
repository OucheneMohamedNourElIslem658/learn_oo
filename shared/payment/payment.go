package payment

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OucheneMohamedNourElIslem658/learn_oo/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type Payment struct {
	instance Config
}

func NewPayment() *Payment {
	return &Payment{
		instance: envs,
	}
}

type Customer struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (p *Payment) CreateCustomer(email, name string) (customer *Customer, err error) {
	instance := p.instance

	user := resty.New()

	requestBody := Customer{
		Email: email,
		Name:  name,
	}

	resp, err := user.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %v", instance.SecretKey)).
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(instance.BaseURL + "/customers")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to create customer")
	}

	var createdCustomer Customer
	if err := json.Unmarshal(resp.Body(), &createdCustomer); err != nil {
		return nil, err
	}

	return &createdCustomer, nil
}

type Product struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	PriceID     string   `json:"price_id,omitempty"`
}

func (p *Payment) CreateProduct(course models.Course) (product *Product, err error) {
	instance := p.instance

	if course.Title == "" {
		return nil, fmt.Errorf("title is required")
	}

	if course.Price <= 0 {
		return nil, fmt.Errorf("price is required")
	}

	user := resty.New()

	requestBody := gin.H{
		"name":        course.Title,
		"description": course.Description,
	}

	images := []string{course.Image.URL}
	if len(images) != 0 {
		requestBody["images"] = images
	}

	resp, err := user.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %v", instance.SecretKey)).
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(instance.BaseURL + "/products")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to create product")
	}

	var createdProduct Product
	if err := json.Unmarshal(resp.Body(), &createdProduct); err != nil {
		return nil, err
	}

	price, err := p.createPrice(
		int(course.Price),
		"dzd",
		createdProduct.ID,
	)

	if err != nil {
		return nil, err
	}

	createdProduct.PriceID = price.ID

	return &createdProduct, nil
}

type Price struct {
	ID        string `json:"id,omitempty"`
	ProductID string `json:"product_id"`
	Amount    int    `json:"amount"`
	Currency  string `json:"currency"`
}

func (p *Payment) createPrice(amount int, currency, productID string) (price *Price, err error) {
	instance := p.instance

	if currency == "" {
		return nil, fmt.Errorf("title is required")
	}

	if amount <= 0 {
		return nil, fmt.Errorf("amount is required")
	}

	if productID == "" {
		return nil, fmt.Errorf("product is required")
	}

	user := resty.New()

	requestBody := Price{
		ProductID: productID,
		Amount:    amount,
		Currency:  currency,
	}

	resp, err := user.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %v", instance.SecretKey)).
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(instance.BaseURL + "/prices")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf(resp.String())
	}

	var createdPrice Price
	if err := json.Unmarshal(resp.Body(), &createdPrice); err != nil {
		return nil, err
	}

	return &createdPrice, nil
}

type Checkout struct {
	ID          string `json:"id,omitempty"`
	CheckoutURL string `json:"checkout_url"`
}

func (p *Payment) MakePayment(successURL, failureURL string, course models.Course) (checkout *Checkout, err error) {
	instance := p.instance

	user := resty.New()

	requestBody := gin.H{
		"success_url": successURL,
		"failure_url": failureURL,
	}

	var items []gin.H
	if course.PaymentPriceID != nil {
		items = append(items, gin.H{
			"price":    *course.PaymentPriceID,
			"quantity": 1,
		})
	} else {
		return nil, fmt.Errorf("course does not have payment price id")
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("courses must be at least more than 1")
	}

	requestBody["items"] = items

	requestBytes, _ := json.MarshalIndent(&requestBody, "\t", "")
	fmt.Println(string(requestBytes))

	resp, err := user.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %v", instance.SecretKey)).
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(instance.BaseURL + "/checkouts")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		fmt.Println(resp.StatusCode())
		var result gin.H
		json.Unmarshal(resp.Body(), &result)
		fmt.Println(result)
		return nil, fmt.Errorf("failed to create checkout")
	}

	var createdCheckout Checkout
	if err := json.Unmarshal(resp.Body(), &createdCheckout); err != nil {
		return nil, err
	}

	return &createdCheckout, nil
}
