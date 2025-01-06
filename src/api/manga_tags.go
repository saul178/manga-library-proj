// this is to make api calls to get mangas by searching based on tags etc.
package api

import "github.com/google/uuid"

type Tags struct {
	ID           uuid.UUID         `json:"id"`
	Type         string            `json:"type"`
	Attributes   TagAttributes     `json:"attributes"`
	Relationship []TagRelationship `json:"relationship"`
	Limit        int               `json:"limit"`
	Offset       int               `json:"offset"`
	Total        int               `json:"total"`
}

type TagAttributes struct {
	Name        map[string]string `json:"name"`
	Description map[string]string `json:"description"`
	Group       string            `json:"group"`
	Version     int               `json:"version"`
}

type TagRelationship []struct {
	ID         uuid.UUID `json:"id"`
	Type       string    `json:"type"`
	Related    string    `json:"related"`
	Attributes struct{}  `json:"attributes"`
}
