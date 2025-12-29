package mollie

import (
	"encoding/json"
	"fmt"
)

// GetCustomer gets a single customer.
func (c *Client) GetCustomer(id string) (*Customer, error) {
	_, respBodyJSON, err := c.request("GET", fmt.Sprintf("/customers/%s", id), nil)

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
