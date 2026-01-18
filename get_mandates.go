package mollie

import (
	"encoding/json"
	"fmt"
)

// GetMandates gets a paginated list of mandates for a customer.
func (c *Client) GetMandates(customerID string) (*Mandates, error) {
	resp, respBodyJSON, err := c.request(
		"GET",
		fmt.Sprintf("/customers/%s/mandates", customerID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	var respBody Mandates
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return &respBody, nil
}
