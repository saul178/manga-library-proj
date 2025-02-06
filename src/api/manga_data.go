package api

import "github.com/google/uuid"

// NOTE: to get specific volumes/chapters i think i have to create another struct that only handles getting the volumes
// and chapters. For the hopes of not bloating this file, i think i should create separate api helpers etc

// NOTE: i might have to refactor this in the future, I'm just not sure how i want to handle some of this info

// NOTE: use Go's memory profiling tool "pprof" for finding bottlenecks in memory usage.

const (
	// endpoint to list all manga related to what is being searched.
	listMangaEndpoint = "manga/"
	// endpoint to get a specific manga by its ID
	getSpecificMangaEndpoint = "/manga/%s"
)

type MangaResponse struct {
	Result   string      `json:"result"`
	Response string      `json:"response"`
	Data     []MangaData `json:"data"`
	Limit    int         `json:"limit"`
	Offset   int         `json:"offset"`
	Total    int         `json:"total"`
}

type MangaData struct {
	ID            uuid.UUID            `json:"id"`
	Type          string               `json:"type"`
	Attributes    MangaAttributes      `json:"attributes"`
	Relationships []MangaRelationships `json:"relationships"`
}

type MangaAttributes struct {
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
	Tags                           []TagsData          `json:"tags"`
	ChapterNumbersResetOnNewVolume bool                `json:"chapterNumbersResetOnNewVolume"`
	AvailableTranslatedLanguages   []string            `json:"availableTranslatedLanguages"`
	LatestUploadedChapter          string              `json:"latestUploadedChapter"`
	State                          string              `json:"state"`
	Version                        int                 `json:"version"`
	CreatedAt                      string              `json:"createdAt"`
	UpdateAt                       string              `json:"updatedAt"`
}

type MangaRelationships struct {
	ID         uuid.UUID `json:"id"`
	Type       string    `json:"type"`
	Related    string    `json:"related"`
	Attributes struct{}  `json:"attributes"`
}

/*
TODO: i want to make a get request to mangadex to get a list of manga to show up, then i want to search for a specific
manga to grab relevant information of that manga.
*/
func (c *MangadexService) listManga() {
}

// place holder functions for now, theyre not finished.
func (m *MangaData) GetMangaID() string {
	return m.ID.String()
}

func (m *MangaData) GetMangaTitle(langCode string) string {
	return m.Attributes.Title[langCode]
}

func (m *MangaData) GetMangaDescriptions(langCode string) string {
	return m.Attributes.Description[langCode]
}
