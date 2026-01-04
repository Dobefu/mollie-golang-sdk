package mollie

import (
	"fmt"
)

// CancelSubscription cancels a subscription.
func (c *Client) CancelSubscription(customerID, subscriptionID string) error {
	_, _, err := c.request(
		"DELETE",
		fmt.Sprintf("/customers/%s/subscriptions/%s", customerID, subscriptionID),
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}
