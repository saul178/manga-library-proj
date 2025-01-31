package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL  = "https://api.mangadex.com"
	coverURL = "https://uploads.mangadex.org"
)

type MangadexService struct {
	client  *http.Client
	header  http.Header
	baseURL string
}

// this is only for guest accounts, if needed i will do an authenticated user if i have to.
func DexClient() *MangadexService {
	client := http.Client{}
	header := http.Header{}
	header.Set("Content-Type", "application/json")

	return &MangadexService{
		client:  &client,
		header:  header,
		baseURL: baseURL,
	}
}

func (c *MangadexService) Request(ctx context.Context, methodCode, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, methodCode, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header = c.header

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	} else if resp.StatusCode != http.StatusOK {
		var er ErrorResponse
		if err = json.NewDecoder(resp.Body).Decode(&er); err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return nil, fmt.Errorf("non 200 status code from mangadex server: %d %s", resp.StatusCode, er.GetErrors())
	}
	return resp, nil
}

// (t any) represents the structs for "type" manga data, author data and cover image data. Hopefully it works the way i think it does
func (c *MangadexService) RequestAndDecodeJson(ctx context.Context, method, url string, body io.Reader, t any) error {
	resp, err := c.Request(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(t); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	return nil
}
