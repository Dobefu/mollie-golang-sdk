package mollie

import (
	"encoding/json"
	"fmt"
)

// GetCustomerSubscriptions gets a paginated list of subscriptions for a customer.
func (c *Client) GetCustomerSubscriptions(customerID string) (*Subscriptions, error) {
	resp, respBodyJSON, err := c.request(
		"GET",
		fmt.Sprintf("/customers/%s/subscriptions", customerID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	var respBody Subscriptions
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return &respBody, nil
}
