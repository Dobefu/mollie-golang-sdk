package mollie

import (
	"encoding/json"
	"fmt"
)

// GetCustomers gets a paginated list of customers.
func (c *Client) GetCustomers() (*Customers, error) {
	resp, respBodyJSON, err := c.request("GET", "/customers", nil)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	var respBody Customers
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUnmarshalRespBody, err)
	}

	return &respBody, nil
}
