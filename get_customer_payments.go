package mollie

import (
	"encoding/json"
	"fmt"
)

// GetCustomerPayments gets a paginated list of customer payments.
func (c *Client) GetCustomerPayments(customerID string) (*Payments, error) {
	_, respBodyJSON, err := c.request(
		"GET",
		fmt.Sprintf("/customers/%s/payments", customerID),
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
