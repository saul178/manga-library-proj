package api

import "github.com/google/uuid"

type AuthorResponse struct {
	Result   string       `json:"result"`
	Response string       `json:"response"`
	Data     []AuthorData `json:"data"`
	Limit    int          `json:"limit"`
	Offset   int          `json:"offset"`
	Total    int          `json:"total"`
}

type AuthorData struct {
	ID            uuid.UUID             `json:"id"`
	Type          string                `json:"type"`
	Attributes    AuthorAttributes      `json:"attributes"`
	Relationships []AuthorRelationships `json:"relationships"`
}

type AuthorAttributes struct {
	Name      string            `json:"name"`
	ImageURL  string            `json:"imageUrl"`
	Biography map[string]string `json:"biography"`
	Twitter   string            `json:"twitter"`
	Pixiv     string            `json:"pixiv"`
	MelonBook string            `json:"melonBook"`
	FanBox    string            `json:"fanBox"`
	Booth     string            `json:"booth"`
	NicoVideo string            `json:"nicoVideo"`
	Skeb      string            `json:"skeb"`
	Fantia    string            `json:"fantia"`
	Tumblr    string            `json:"tumblr"`
	Youtube   string            `json:"youtube"`
	Weibo     string            `json:"weibo"`
	Naver     string            `json:"naver"`
	Namicomi  string            `json:"namicomi"`
	Website   string            `json:"website"`
	Version   int               `json:"version"`
	CreatedAt string            `json:"createdAt"`
	UpdatedAt string            `json:"updatedAt"`
}

type AuthorRelationships struct {
	ID         uuid.UUID `json:"id"`
	Type       string    `json:"type"`
	Related    string    `json:"related"`
	Attributes struct{}  `json:"attributes"`
}
