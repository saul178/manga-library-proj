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
			TagID         string `json:"id"`
			TagType       string `json:"type"`
			TagAttributes struct {
				TagName struct {
					TagProperty1 string `json:"property1"`
					TagProperty2 string `json:"property2"`
				}
			}
		}
	}
}
