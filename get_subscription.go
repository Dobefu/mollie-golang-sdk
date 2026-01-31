package mollie

import (
	"encoding/json"
	"fmt"
)

// GetSubscription gets a single subscription.
func (c *Client) GetSubscription(
	customerID,
	subscriptionID string,
) (*Subscription, error) {
	resp, respBodyJSON, err := c.request(
		"GET",
		fmt.Sprintf("/customers/%s/subscriptions/%s", customerID, subscriptionID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	var respBody Subscription
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUnmarshalRespBody, err)
	}

	return &respBody, nil
}
