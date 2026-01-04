package mollie

import (
	"encoding/json"
	"fmt"
)

// GetSubscriptionPayments gets a paginated list of subscription payments.
func (c *Client) GetSubscriptionPayments(
	customerID string,
	subscriptionID string,
) (*Payments, error) {
	_, respBodyJSON, err := c.request(
		"GET",
		fmt.Sprintf("/customers/%s/subscriptions/%s", customerID, subscriptionID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	var respBody Payments
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return &respBody, nil
}
