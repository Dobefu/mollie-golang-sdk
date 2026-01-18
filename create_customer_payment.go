package mollie

import (
	"encoding/json"
	"fmt"
)

// CreateCustomerPaymentBody represents a single CreateCustomerPayment body.
type CreateCustomerPaymentBody struct {
	Description string            `json:"description" url:"description"`
	Amount      Amount            `json:"amount" url:"amount"`
	RedirectURL string            `json:"redirectUrl,omitempty" url:"redirectUrl,omitempty"`
	WebhookURL  string            `json:"webhookUrl,omitempty" url:"webhookUrl,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty" url:"metadata,omitempty"`
}

// CreateCustomerPayment creates a new customer payment.
func (c *Client) CreateCustomerPayment(
	customerID string,
	body CreateCustomerPaymentBody,
) (*Payment, error) {
	resp, respBodyJSON, err := c.request(
		"POST",
		fmt.Sprintf("/customers/%s/payments", customerID),
		body,
	)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	var respBody Payment
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return &respBody, nil
}
