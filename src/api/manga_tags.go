// this is to make api calls to get mangas by searching based on tags etc.
package api

import "github.com/google/uuid"

type TagsList struct {
	Result   string     `json:"result"`
	Response string     `json:"response"`
	Data     []TagsData `json:"data"`
	Limit    int        `json:"limit"`
	Offset   int        `json:"offset"`
	Total    int        `json:"total"`
}

type TagsData struct {
	ID            uuid.UUID          `json:"id"`
	Type          string             `json:"type"`
	Attributes    TagAttributes      `json:"attributes"`
	Relationships []TagRelationships `json:"relationships"`
}

type TagAttributes struct {
	Name        map[string]string `json:"name"`
	Description map[string]string `json:"description"`
	Group       string            `json:"group"`
	Version     int               `json:"version"`
}

type TagRelationships []struct {
	ID         uuid.UUID `json:"id"`
	Type       string    `json:"type"`
	Related    string    `json:"related"`
	Attributes struct{}  `json:"attributes"`
}
