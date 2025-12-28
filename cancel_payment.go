package mollie

import (
	"encoding/json"
	"fmt"
)

// CancelPayment cancels a payment.
func (c *Client) CancelPayment(id string) error {
	_, respBodyJSON, err := c.request("DELETE", fmt.Sprintf("/payments/%s", id), nil)

	if err != nil {
		return err
	}

	var respBody Payment
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return nil
}
