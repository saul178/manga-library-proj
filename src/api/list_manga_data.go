package api

// NOTE: to get specific volumes/chapters i think i have to create another struct that only handles getting the volumes
// and chapters. for the hopes of not bloating this file, i think i should create seperate api helpers etc

// NOTE: i might have to refactor this in the future, im just not sure how i want to handle some of this info
type MangaData struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Title                          map[string]string   `json:"title"`
			AltTitles                      []map[string]string `json:"altTitles"`
			Description                    map[string]string   `json:"description"`
			IsLocked                       bool                `json:"isLocked"`
			Links                          map[string]string   `json:"links"`
			OriginalLanguage               string              `json:"originalLanguage"`
			LastVolume                     string              `json:"lastVolume"`
			LastChapter                    string              `json:"lastChapter"`
			PublicationDemographic         string              `json:"publicationDemographic"`
			Status                         string              `json:"status"`
			Year                           int                 `json:"year"`
			ContentRating                  string              `json:"contentRating"`
			ChapterNumbersResetOnNewVolume bool                `json:"chapterNumbersResetOnNewVolume"`
			AvailableTranslatedLanguages   []string            `json:"availableTranslatedLanguages"`
			LatestUploadedChapter          string              `json:"latestUploadedChapter"`

			Tags []struct {
				ID         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Name        map[string]string `json:"name"`
					Description map[string]string `json:"description"`
				}
			} `json:"tags"`
		} `json:"attributes"`
	} `json:"data"`
}
