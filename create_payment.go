package mollie

import (
	"encoding/json"
	"fmt"
)

// CreatePaymentBody represents a single CreatePayment body.
type CreatePaymentBody struct {
	CustomerID   string            `json:"customerId,omitempty" url:"customerId,omitempty"`
	SequenceType string            `json:"sequenceType,omitempty" url:"sequenceType,omitempty"`
	Description  string            `json:"description" url:"description"`
	Amount       Amount            `json:"amount" url:"amount"`
	RedirectURL  string            `json:"redirectUrl,omitempty" url:"redirectUrl,omitempty"`
	WebhookURL   string            `json:"webhookUrl,omitempty" url:"webhookUrl,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty" url:"metadata,omitempty"`
}

// CreatePayment creates a new payment.
func (c *Client) CreatePayment(body CreatePaymentBody) (*Payment, error) {
	_, respBodyJSON, err := c.request("POST", "/payments", body)

	if err != nil {
		return nil, err
	}

	var respBody *Payment
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	redirectURL := fmt.Sprintf(
		"%s?payment_id=%s",
		respBody.RedirectURL,
		respBody.ID,
	)

	respBody, err = c.UpdatePayment(respBody.ID, UpdatePaymentBody{
		Description: body.Description,
		RedirectURL: redirectURL,
		WebhookURL:  body.WebhookURL,
		Metadata:    body.Metadata,
	})

	if err != nil {
		return nil, fmt.Errorf("could not update payment: %s", err.Error())
	}

	return respBody, nil
}
