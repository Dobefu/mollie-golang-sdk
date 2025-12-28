package mollie

import (
	"encoding/json"
	"fmt"
)

// GetPayments gets a paginated list of payments.
func (c *Client) GetPayments() error {
	_, respBodyJSON, err := c.request("GET", "/payments", nil)

	if err != nil {
		return err
	}

	var respBody struct {
		Count    int `json:"count" url:"count"`
		Embedded struct {
			Payments []Payment `json:"payments" url:"payments"`
		} `json:"_embedded" url:"_embedded"`
	}

	err = json.Unmarshal(respBodyJSON, &respBody)

	if err != nil {
		return fmt.Errorf("could not unmarshal response body: %s", err.Error())
	}

	return nil
}
