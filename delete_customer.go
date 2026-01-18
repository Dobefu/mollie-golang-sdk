package mollie

import (
	"fmt"
)

// DeleteCustomer deletes a customer.
func (c *Client) DeleteCustomer(id string) error {
	resp, _, err := c.request("DELETE", fmt.Sprintf("/customers/%s", id), nil)

	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	return nil
}
