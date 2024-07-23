package helpers

type MangaDexResponse struct {
	Result   string `json:"result"`
	Response string `json:"response"`
}

type MangaData struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Title     map[string]string   `json:"title"`
			AltTitles []map[string]string `json:"altTitles"`
		} `json:"attributes"`
	} `json:"data"`
}
