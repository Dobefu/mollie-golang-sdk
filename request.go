package mollie

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	// ErrMarshalReqBody specifies when a request body could not be marshalled.
	ErrMarshalReqBody = errors.New("could not marshal request body")

	// ErrSendRequest specifies when a request could not be sent.
	ErrSendRequest = errors.New("could not send request")

	// ErrRequestFailed specifies when a request fails.
	ErrRequestFailed = errors.New("request failed")
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
			return nil, nil, fmt.Errorf("%w: %w", ErrMarshalReqBody, err)
		}

		bodyBuffer = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", apiURL, path),
		bodyBuffer,
	)

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", ErrSendRequest, err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	resp, err := c.httpClient.Do(req)
	defer func() { _ = resp.Body.Close() }()

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", ErrRequestFailed, err)
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
