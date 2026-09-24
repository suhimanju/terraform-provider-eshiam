// Package client provides a thin, generic OAuth2 client-credentials REST client.
//
// It centralises authentication and request handling so resources never touch
// transport concerns directly. The token is cached and refreshed automatically
// when it expires. Replace the request/response helpers with a generated SDK if
// you have one.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golang.org/x/oauth2"
)

// Config holds the settings required to construct an APIClient.
type Config struct {
	BaseURL      string
	ClientId     string
	ClientSecret string
	TokenURL     string
	// HTTPClient is optional; a sensible default is used when nil.
	HTTPClient *http.Client
}

// APIClient is an authenticated gateway to the target REST API.
type APIClient struct {
	config     Config
	httpClient *http.Client

	mu    sync.Mutex
	token *oauth2.Token
}

// NewAPIClient builds an APIClient from the given configuration.
func NewAPIClient(config Config) *APIClient {
	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 30 * time.Second}
	}
	return &APIClient{config: config, httpClient: httpClient}
}

// BaseURL returns the API base URL.
func (c *APIClient) BaseURL() string { return c.config.BaseURL }

// CreateObject POSTs body to path and returns the decoded JSON object.
func (c *APIClient) CreateObject(ctx context.Context, path string, body any) (map[string]any, *http.Response, error) {
	var out map[string]any
	resp, err := c.Do(ctx, http.MethodPost, path, body, &out)
	return out, resp, err
}

// GetObject GETs path and returns the decoded JSON object.
func (c *APIClient) GetObject(ctx context.Context, path string) (map[string]any, *http.Response, error) {
	var out map[string]any
	resp, err := c.Do(ctx, http.MethodGet, path, nil, &out)
	return out, resp, err
}

// ListObjects GETs path and returns the decoded JSON array.
func (c *APIClient) ListObjects(ctx context.Context, path string) ([]map[string]any, *http.Response, error) {
	var out []map[string]any
	resp, err := c.Do(ctx, http.MethodGet, path, nil, &out)
	return out, resp, err
}

// UpdateObject sends body to path using the given method (PUT or PATCH) and
// returns the decoded JSON object.
func (c *APIClient) UpdateObject(ctx context.Context, method, path string, body any) (map[string]any, *http.Response, error) {
	var out map[string]any
	resp, err := c.Do(ctx, method, path, body, &out)
	return out, resp, err
}

// DeleteObject DELETEs path.
func (c *APIClient) DeleteObject(ctx context.Context, path string) (*http.Response, error) {
	return c.Do(ctx, http.MethodDelete, path, nil, nil)
}

// Do performs an authenticated request against uri (relative to BaseURL).
//
// When body is non-nil it is JSON-encoded. On a 2xx response, and when out is
// non-nil, the response body is decoded into out.
func (c *APIClient) Do(ctx context.Context, method, uri string, body, out any) (*http.Response, error) {
	var reader io.Reader
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("encoding request body: %w", err)
		}
		reader = bytes.NewReader(encoded)
	}

	fullURL, err := url.Parse(c.config.BaseURL + uri)
	if err != nil {
		return nil, fmt.Errorf("parsing request URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL.String(), reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	token, err := c.authToken(ctx)
	if err != nil {
		return nil, err
	}
	token.SetAuthHeader(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return resp, fmt.Errorf("error calling %s %s: %s", method, fullURL.String(), resp.Status)
	}

	if out != nil {
		if err := decode(resp, out); err != nil {
			return resp, err
		}
	}
	return resp, nil
}

// decode reads and JSON-decodes the response body into out, leaving the body
// readable again for callers that inspect it (e.g. for error reporting).
func decode(resp *http.Response, out any) error {
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	resp.Body = io.NopCloser(bytes.NewBuffer(raw))
	if err != nil {
		return err
	}
	if len(raw) == 0 {
		return nil
	}
	return json.Unmarshal(raw, out)
}

// authToken returns a cached, valid token or fetches a fresh one.
func (c *APIClient) authToken(ctx context.Context) (*oauth2.Token, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.token != nil && c.token.Valid() {
		return c.token, nil
	}

	token, err := c.fetchToken(ctx)
	if err != nil {
		return nil, err
	}
	c.token = token
	return token, nil
}

// fetchToken performs the OAuth2 client-credentials grant.
func (c *APIClient) fetchToken(ctx context.Context) (*oauth2.Token, error) {
	tflog.Info(ctx, "Requesting access token from "+c.config.TokenURL)

	form := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {c.config.ClientId},
		"client_secret": {c.config.ClientSecret},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.config.TokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("token request failed: %s: %s", resp.Status, string(raw))
	}

	var token oauth2.Token
	if err := json.Unmarshal(raw, &token); err != nil {
		return nil, err
	}
	return &token, nil
}
