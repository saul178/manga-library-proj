package api

import "github.com/google/uuid"

type CoverResponse struct {
	Result   string      `json:"result"`
	Response string      `json:"response"`
	Data     []CoverData `json:"data"`
	Limit    int         `json:"limit"`
	Offset   int         `json:"offset"`
	Total    int         `json:"total"`
}

type CoverData struct {
	ID            uuid.UUID            `json:"id"`
	Type          string               `json:"type"`
	Attributes    CoverAttributes      `json:"attributes"`
	Relationships []CoverRelationships `json:"relationships"`
}

type CoverAttributes struct {
	Volume      string `json:"volume"`
	FileName    string `json:"fileName"`
	Description string `json:"description"`
	Locale      string `json:"locale"`
	Version     int    `json:"version"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type CoverRelationships struct {
	ID         uuid.UUID `json:"id"`
	Type       string    `json:"type"`
	Related    string    `json:"related"`
	Attributes struct{}  `json:"attributes"`
}
