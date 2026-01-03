package mollie

import (
	"encoding/json"
	"fmt"
)

// GetMandate gets a single mandate.
func (c *Client) GetMandate(customerID, mandateID string) (*Mandate, error) {
	_, respBodyJSON, err := c.request(
		"GET",
		fmt.Sprintf("/customers/%s/mandates/%s", customerID, mandateID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	var respBody Mandate
	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return &respBody, nil
}
