package mollie

import (
	"encoding/json"
	"fmt"
)

// GetSubscriptions gets a paginated list of subscriptions for a customer.
func (c *Client) GetSubscriptions() (*Subscriptions, error) {
	_, respBodyJSON, err := c.request("GET", "/subscriptions", nil)

	if err != nil {
		return nil, err
	}

	var respBody Subscriptions
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return &respBody, nil
}
