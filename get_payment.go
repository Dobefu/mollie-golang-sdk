package mollie

import (
	"encoding/json"
	"fmt"
)

// GetPayment gets a single payment.
func (c *Client) GetPayment(id string) error {
	_, respBodyJSON, err := c.request("GET", fmt.Sprintf("/payments/%s", id), nil)

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
