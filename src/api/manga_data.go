package api

// NOTE: to get specific volumes/chapters i think i have to create another struct that only handles getting the volumes
// and chapters. for the hopes of not bloating this file, i think i should create seperate api helpers etc
type MangaDexResponse struct {
	Result   string `json:"result"`
	Response string `json:"response"`
}

type MangaData struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Title       map[string]string   `json:"title"`
			AltTitles   []map[string]string `json:"altTitles"`
			Description map[string]string   `json:"description"`
			IsLocked    bool                `json:"isLocked"`
			Links       map[string]string   `json:"links"`
		} `json:"attributes"`
	} `json:"data"`
}
