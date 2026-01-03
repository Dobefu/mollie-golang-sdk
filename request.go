package mollie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// RequestBody specifies a request body.
type RequestBody any

func (c *Client) request(
	method string,
	path string,
	body RequestBody,
) (*http.Response, []byte, error) {
	var bodyBuffer io.Reader

	if body != nil {
		jsonBody, err := json.Marshal(body)

		if err != nil {
			return nil, nil, fmt.Errorf("could marshal request body: %s", err.Error())
		}

		bodyBuffer = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", apiURL, path),
		bodyBuffer,
	)

	if err != nil {
		return nil, nil, fmt.Errorf("could not send request: %s", err.Error())
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	resp, err := c.httpClient.Do(req)
	defer func() { _ = resp.Body.Close() }()

	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %s", err.Error())
	}

	respBodyJSON, err := io.ReadAll(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var respBody Error

		err = json.Unmarshal(respBodyJSON, &respBody)

		if err != nil {
			return nil, nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
		}

		return nil, nil, fmt.Errorf(
			"request failed with status code %d: %s",
			resp.StatusCode,
			respBody.Detail,
		)
	}

	if err != nil {
		return nil, nil, fmt.Errorf("could read response body: %s", err.Error())
	}

	return resp, respBodyJSON, nil
}
