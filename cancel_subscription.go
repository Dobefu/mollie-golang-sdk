package mollie

import (
	"fmt"
)

// CancelSubscription cancels a subscription.
func (c *Client) CancelSubscription(customerID, subscriptionID string) error {
	resp, _, err := c.request(
		"DELETE",
		fmt.Sprintf("/customers/%s/subscriptions/%s", customerID, subscriptionID),
		nil,
	)

	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	return nil
}
