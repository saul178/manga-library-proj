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
			Title struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			}
		}
		Tags []struct {
			ID         string `json:"id"`
			Type       string `json:"type"`
			Attributes struct {
				Name struct {
					Property1 string `json:"property1"`
					Property2 string `json:"property2"`
				}
			}
		}
	}
}
