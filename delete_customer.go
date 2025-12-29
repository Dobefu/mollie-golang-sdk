package mollie

import (
	"fmt"
)

// DeleteCustomer deletes a customer.
func (c *Client) DeleteCustomer(id string) error {
	_, _, err := c.request("DELETE", fmt.Sprintf("/customers/%s", id), nil)

	if err != nil {
		return err
	}

	return nil
}
