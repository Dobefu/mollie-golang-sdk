package mollie

import (
	"encoding/json"
	"fmt"
)

// GetCustomers gets a paginated list of customers.
func (c *Client) GetCustomers() (*Customers, error) {
	_, respBodyJSON, err := c.request("GET", "/customers", nil)

	if err != nil {
		return nil, err
	}

	var respBody Customers
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return &respBody, nil
}
