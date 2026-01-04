package mollie

import (
	"encoding/json"
	"fmt"
)

// UpdateSubscriptionBody represents a single UpdateSubscription body.
type UpdateSubscriptionBody struct {
	Description string            `json:"description,omitempty" url:"description,omitempty"`
	Interval    string            `json:"interval,omitempty" url:"interval,omitempty"`
	Times       int               `json:"times,omitempty" url:"times,omitempty"`
	WebhookURL  string            `json:"webhookUrl,omitempty" url:"webhookUrl,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty" url:"metadata,omitempty"`
}

// UpdateSubscription updates an existing subscription.
func (c *Client) UpdateSubscription(
	customerID string,
	subscriptionID string,
	body UpdateSubscriptionBody,
) (*Subscription, error) {
	_, respBodyJSON, err := c.request(
		"PATCH",
		fmt.Sprintf("/customers/%s/subscriptions/%s", customerID, subscriptionID),
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
