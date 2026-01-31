package mollie

import (
	"encoding/json"
	"fmt"
)

// GetPayments gets a paginated list of payments.
func (c *Client) GetPayments() (*Payments, error) {
	resp, respBodyJSON, err := c.request("GET", "/payments", nil)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	var respBody Payments
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUnmarshalRespBody, err)
	}

	return &respBody, nil
}
