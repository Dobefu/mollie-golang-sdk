package mollie

import (
	"encoding/json"
	"fmt"
)

// CreateSubscriptionBody represents a single CreateSubscription body.
type CreateSubscriptionBody struct {
	MandateID   string `json:"mandateId,omitempty" url:"mandateId,omitempty"`
	Amount      Amount `json:"amount" url:"amount"`
	Interval    string `json:"interval" url:"interval"`
	Description string `json:"description" url:"description"`
	Times       int    `json:"times,omitempty" url:"times,omitempty"`
	StartDate   Date   `json:"startDate" url:"startDate"`
	WebhookURL  string `json:"webhookUrl" url:"webhookUrl"`
}

// CreateSubscription creates a new subscription.
func (c *Client) CreateSubscription(
	customerID string,
	body CreateSubscriptionBody,
) (*Subscription, error) {
	_, respBodyJSON, err := c.request(
		"POST",
		fmt.Sprintf("/customers/%s/subscriptions", customerID),
		body,
	)

	if err != nil {
		return nil, err
	}

	var respBody Subscription
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return &respBody, nil
}
