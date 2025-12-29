package jeleoreversegeo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const defaultBaseURL = "https://slimigeo.jeleo.zone.id"

// Client is a minimal Reverse Geocoding API client.
type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new client with the provided API key and optional base URL.
// If baseURL is empty, the public default is used.
func NewClient(apiKey string, baseURL string) (*Client, error) {
	if strings.TrimSpace(apiKey) == "" {
		return nil, errors.New("apiKey is required")
	}
	if strings.TrimSpace(baseURL) == "" {
		baseURL = defaultBaseURL
	}
	return &Client{
		APIKey:  apiKey,
		BaseURL: strings.TrimRight(baseURL, "/"),
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

// ReverseGeo performs reverse geocoding for the given lat/lng and unmarshals the
// JSON response into v. v can be a pointer to map[string]any or a struct matching
// your expected response schema.
func (c *Client) ReverseGeo(lat, lng float64, v any) error {
	if c == nil {
		return errors.New("client is nil")
	}
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	}
	u, err := url.Parse(c.BaseURL + "/reverse-geo")
	if err != nil {
		return fmt.Errorf("invalid base URL: %w", err)
	}
	q := u.Query()
	q.Set("lat", fmt.Sprintf("%g", lat))
	q.Set("lng", fmt.Sprintf("%g", lng))
	q.Set("api_key", c.APIKey)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("performing request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("reverse-geo failed: %s %s", resp.Status, string(b))
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(v); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}
	return nil
}
