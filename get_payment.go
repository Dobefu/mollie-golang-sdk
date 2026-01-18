package mollie

import (
	"encoding/json"
	"fmt"
)

// GetPayment gets a single payment.
func (c *Client) GetPayment(id string) (*Payment, error) {
	resp, respBodyJSON, err := c.request("GET", fmt.Sprintf("/payments/%s", id), nil)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	var respBody Payment
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return &respBody, nil
}
