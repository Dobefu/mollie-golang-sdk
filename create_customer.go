package mollie

import (
	"encoding/json"
	"fmt"
)

// CreateCustomerBody represents a single CreateCustomer body.
type CreateCustomerBody struct {
	Name  string `json:"name" url:"name"`
	Email string `json:"email" url:"email"`
}

// CreateCustomer creates a new customer.
func (c *Client) CreateCustomer(body CreateCustomerBody) (*Customer, error) {
	_, respBodyJSON, err := c.request("POST", "/customers", body)

	if err != nil {
		return nil, err
	}

	var respBody Customer
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return &respBody, nil
}
