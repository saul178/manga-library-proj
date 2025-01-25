package api

import (
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
