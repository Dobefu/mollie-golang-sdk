package mollie

import (
	"encoding/json"
	"fmt"
)

// UpdatePaymentBody represents a single UpdatePayment body.
type UpdatePaymentBody struct {
	Description string            `json:"description,omitempty" url:"description, omitempty"`
	RedirectURL string            `json:"redirectUrl,omitempty" url:"redirectUrl,omitempty"`
	WebhookURL  string            `json:"webhookUrl,omitempty" url:"webhookUrl,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty" url:"metadata,omitempty"`
}

// UpdatePayment updates an existing payment.
func (c *Client) UpdatePayment(
	paymentID string,
	body UpdatePaymentBody,
) (*Payment, error) {
	_, respBodyJSON, err := c.request(
		"PATCH",
		fmt.Sprintf("/payments/%s", paymentID),
		body,
	)

	if err != nil {
		return nil, err
	}

	var respBody Payment
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return &respBody, nil
}
