package api

import (
	"context"
	"net/http"
)

const baseURL string = "https://api.mangadex.com"

type MangadexService struct {
	client *http.Client
	header http.Header
}

func DexClient() *MangadexService {
}
