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

	// ErrSendReq specifies when a request could not be sent.
	ErrSendReq = errors.New("could not send request")

	// ErrReqFailed specifies when a request fails.
	ErrReqFailed = errors.New("request failed")

	// ErrReqFailedWithStatusCode specifies when a request fails with a status code.
	ErrReqFailedWithStatusCode = errors.New("request failed with status code")

	// ErrReadRespBody specifies when a response body could not be read.
	ErrReadRespBody = errors.New("could not read response body")
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
		return nil, nil, fmt.Errorf("%w: %w", ErrSendReq, err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	resp, err := c.httpClient.Do(req)
	defer func() { _ = resp.Body.Close() }()

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", ErrReqFailed, err)
	}

	respBodyJSON, err := io.ReadAll(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var respBody Error

		err = json.Unmarshal(respBodyJSON, &respBody)

		if err != nil {
			return nil, nil, fmt.Errorf(
				"%w: %d",
				ErrReqFailedWithStatusCode,
				resp.StatusCode,
			)
		}

		return nil, nil, fmt.Errorf(
			"%w: %d: %s",
			ErrReqFailedWithStatusCode,
			resp.StatusCode,
			respBody.Detail,
		)
	}

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", ErrReadRespBody, err)
	}

	return resp, respBodyJSON, nil
}
