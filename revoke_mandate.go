package mollie

import (
	"fmt"
)

// RevokeMandate revokes a single mandate.
func (c *Client) RevokeMandate(customerID, mandateID string) error {
	_, _, err := c.request(
		"DELETE",
		fmt.Sprintf("/customers/%s/mandates/%s", customerID, mandateID),
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}
