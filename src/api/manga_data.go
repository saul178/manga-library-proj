package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

// NOTE: use Go's memory profiling tool "pprof" for finding bottlenecks in memory usage.

const (
	// endpoint to list all manga related to what is being searched.
	listMangaEndpoint = "/manga"
	// endpoint to get a specific manga by its ID
	getSpecificMangaEndpoint = "/manga/%s"
)

// TODO: omit empty fields from some of the responses that might be not have data for all structs in this project.
type MangaResponse struct {
	Result   string `json:"result"`
	Response string `json:"response"`
}

type SingleMangaResponse struct {
	MangaResponse
	Data MangaData `json:"data"`
}

type MultiMangaResponse struct {
	MangaResponse
	ListMangaData []MangaData `json:"data"`
	Limit         int         `json:"limit"`
	Offset        int         `json:"offset"`
	Total         int         `json:"total"`
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
	ID         uuid.UUID   `json:"id"`
	Type       string      `json:"type"`
	Related    string      `json:"related,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

/*
TODO: i want to make a get request to mangadex to get a list of manga to show up, then i want to search for a specific
manga to grab relevant information of that manga.
*/

/*
NOTE:
when searching for manga the url parameters that it accepts are titles of the manga and the limit
the context that should be passed down is context.WithTimeout so when mangadex hangs we don't continue to send the
request and instead handle that gracefully.
*/
func (s *MangadexService) SearchMangas(ctx context.Context, params url.Values) (*MultiMangaResponse, error) {
	u, err := url.Parse(s.baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}
	u.Path = listMangaEndpoint
	u.RawQuery = params.Encode()

	var mangaListResp MultiMangaResponse
	err = s.RequestAndDecodeJson(ctx, http.MethodGet, u.String(), nil, &mangaListResp)
	return &mangaListResp, err
}

/*
NOTE:
fetches a single manga by its ID, valid paramaters that can be included are: manga, cover_art, author,
artist, tag, and creator.
*/

func (s *MangadexService) GetManga(ctx context.Context, id uuid.UUID, params url.Values) (*SingleMangaResponse, error) {
	u, err := url.Parse(s.baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}
	u.Path = fmt.Sprintf(getSpecificMangaEndpoint, id)
	u.RawQuery = params.Encode()
	fmt.Println(u)

	var manga SingleMangaResponse
	err = s.RequestAndDecodeJson(ctx, http.MethodGet, u.String(), nil, &manga)
	return &manga, err
}

// TODO: place holder functions for now, theyre not finished. maybe have a custom struct with the data i care about?
// maybe have this in it's own file since it handles a different job of just getting specific info and isnt in charge of
// doing api calls?

type MangaInfo struct {
	Title         map[string]string
	Description   map[string]string
	Links         map[string]string
	LastVolume    string
	LastChapter   string
	Status        string
	Year          int
	ContentRating string
	Tags          []TagsData
	CreatedAt     string
	UpdateAt      string
}

func (m *MangaData) GetMangaID() string {
	return m.ID.String()
}

func (m *MangaData) GetMangaTitle(langCode string) string {
	return m.Attributes.Title[langCode]
}

func (m *MangaData) GetMangaDescriptions(langCode string) string {
	return m.Attributes.Description[langCode]
}
