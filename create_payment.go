package mollie

import (
	"encoding/json"
	"fmt"
)

// CreatePaymentBody represents a single CreatePayment body.
type CreatePaymentBody struct {
	Description string            `json:"description" url:"description"`
	Amount      Amount            `json:"amount" url:"amount"`
	RedirectURL string            `json:"redirectUrl,omitempty" url:"redirectUrl,omitempty"`
	WebhookURL  string            `json:"webhookUrl,omitempty" url:"webhookUrl,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty" url:"metadata,omitempty"`
}

// CreatePayment creates a new payment.
func (c *Client) CreatePayment(body CreatePaymentBody) error {
	_, respBodyJSON, err := c.request("POST", "/payments", body)

	if err != nil {
		return err
	}

	var respBody Payment
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return nil
}
