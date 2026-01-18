package mollie

import (
	"fmt"
)

// CancelPayment cancels a payment.
func (c *Client) CancelPayment(id string) error {
	resp, _, err := c.request("DELETE", fmt.Sprintf("/payments/%s", id), nil)

	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	return nil
}
