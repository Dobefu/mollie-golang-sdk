package mollie

import (
	"fmt"
)

// CancelPayment cancels a payment.
func (c *Client) CancelPayment(id string) error {
	_, _, err := c.request("DELETE", fmt.Sprintf("/payments/%s", id), nil)

	if err != nil {
		return err
	}

	return nil
}
