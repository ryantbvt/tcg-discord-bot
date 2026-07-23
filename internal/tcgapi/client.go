package tcgapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://api.tcgdex.net/v2"

type Client struct {
	httpClient *http.Client
	lang       string
}

func NewClient(lang string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		lang:       lang,
	}
}

func (c *Client) get(path string, v any) error {
	url := fmt.Sprintf("%s/%s/%s", baseURL, c.lang, path)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("request to %s failed: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return ErrNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d from %s", resp.StatusCode, url)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
