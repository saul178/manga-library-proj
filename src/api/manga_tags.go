// this is to make api calls to get mangas by searching based on tags etc.
package api

import "github.com/google/uuid"

type Tags struct {
	Data []struct {
		ID   uuid.UUID `json:"id"`
		Type string    `json:"type"`

		Attributes struct {
			Name        map[string]string `json:"name"`
			Description map[string]string `json:"description"`
			Group       string            `json:"group"`
			Version     int               `json:"version"`
		} `json:"attributes"`

		Relationship []struct {
			ID         uuid.UUID `json:"id"`
			Type       string    `json:"type"`
			Related    string    `json:"related"`
			Attributes struct{}  `json:"attributes"`
		}

		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	} `json:"data"`
}
