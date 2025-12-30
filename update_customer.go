package mollie

import (
	"encoding/json"
	"fmt"
)

// UpdateCustomerBody represents a single UpdateCustomer body.
type UpdateCustomerBody struct {
	Name  string `json:"name" url:"name"`
	Email string `json:"email" url:"email"`
}

// UpdateCustomer updates an existing customer.
func (c *Client) UpdateCustomer(id string, body UpdateCustomerBody) (*Customer, error) {
	_, respBodyJSON, err := c.request(
		"PATCH",
		fmt.Sprintf("/customers/%s", id),
		body,
	)

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
