package mollie

import "net/http"

// Client specifies a single Mollie client.
type Client struct {
	APIKey string
	Mode   Mode

	httpClient *http.Client
}

// NewClient creates a new Mollie client.
func NewClient(apiKey string, mode Mode) *Client {
	c := &Client{
		APIKey: apiKey,
		Mode:   mode,

		httpClient: http.DefaultClient,
	}

	return c
}
