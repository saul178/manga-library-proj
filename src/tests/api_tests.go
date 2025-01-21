package tests

/* TODO:
   1.) learn why the relationships array struct comes out blank.
	- the relationships array is pointing to other manga that have a relation with the series you looked up
	- and it will show what type of manga it has a relation with. Negima doujinshi -> to Negima series etc
   2.) figure out how to separate these into fuctions
   3.) really learn these standard libraries
*/
import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/saul178/manga-library-proj/src/api"
)

const baseUrl = "https://api.mangadex.org"

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func TestClient() *Client {
	return &Client{
		HTTPClient: &http.Client{},
		BaseURL:    baseUrl,
	}
}

func (c *Client) getCoverArt(mangaTitle string, limit int) (api.CoverData, error) {
	endpoint := fmt.Sprintf("%s/covers", c.BaseURL)
	getManga, err := c.SearchManga(mangaTitle, limit)
	if err != nil {
		return api.CoverData{}, err
	}

	mangaID := getManga[0].ID
	fmt.Println(getManga[0].Attributes.Title, mangaID)

	params := url.Values{}
	params.Add("mangaID", mangaID.String())

	req, err := http.Get(endpoint)
	if err != nil {
		return api.CoverData{}, err
	}
	defer req.Body.Close()

	var coverResp api.CoverResponse

	return api.CoverData{}, nil
}

func (c *Client) SearchAuthors(name string, limit int) ([]api.AuthorData, error) {
	endpoint := fmt.Sprintf("%s/author", c.BaseURL)
	params := url.Values{}
	params.Add("name", name)
	params.Add("limit", fmt.Sprintf("%d", limit))

	req, err := http.NewRequest("GET", endpoint+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var authorResp api.AuthorResponse
	err = json.NewDecoder(resp.Body).Decode(&authorResp)
	if err != nil {
		return nil, err
	}

	var authorList []api.AuthorData
	for _, a := range authorResp.Data {
		authorList = append(authorList, a)
	}

	return authorList, nil
}

func (c *Client) SearchManga(title string, limit int) ([]api.MangaData, error) {
	endpoint := fmt.Sprintf("%s/manga", c.BaseURL)
	params := url.Values{}
	params.Add("title", title)
	params.Add("limit", fmt.Sprintf("%d", limit))

	req, err := http.NewRequest("GET", endpoint+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var result api.Manga
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	var mangaList []api.MangaData

	for _, item := range result.Data {
		mangaList = append(mangaList, item)
	}

	return mangaList, nil
}

// TODO: need to test retreiving manga volume and chapter information
// NOTE: to get manga volumes & chapter info i need a specific manga id that is searched for.
// func fetchMangaID(mangaTitle string)
func (c *Client) GetMangaVolumesInfo(mangaTitle string) (api.MangaVolumeResponse, error) {
	getManga, err := c.SearchManga(mangaTitle, 1)
	if err != nil {
		return api.MangaVolumeResponse{}, errors.New("need to learn how to error handle correctly.")
	}

	mangaID := getManga[0].ID
	fmt.Println(getManga[0].Attributes.Title, mangaID)

	endpoint := fmt.Sprintf("%s/manga/%s/aggregate", c.BaseURL, mangaID)
	params := url.Values{}
	params.Add("mangaID", mangaID.String())

	resp, err := http.Get(endpoint)
	if err != nil {
		return api.MangaVolumeResponse{}, errors.New("bad request \n")
	}
	defer resp.Body.Close()

	var mangaVolResp api.MangaVolumeResponse
	if err := json.NewDecoder(resp.Body).Decode(&mangaVolResp); err != nil {
		return api.MangaVolumeResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}
	test, _ := json.MarshalIndent(mangaVolResp, "", " ")
	fmt.Println(string(test))

	return mangaVolResp, nil
}

func (c *Client) SearchByTags(includedTags, excludedTags []string, limit int) ([]api.MangaData, error) {
	endpoint := fmt.Sprintf("%s/manga/tag", c.BaseURL)
	tagsResp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer tagsResp.Body.Close()

	var tags api.TagsList
	tagErr := json.NewDecoder(tagsResp.Body).Decode(&tags)
	if tagErr != nil {
		return nil, err
	}

	includedTagsIDs, excludedTagsIDs := extractTagIds(tags, includedTags, excludedTags)

	params := url.Values{}
	for _, tagID := range includedTagsIDs {
		params.Add("includedTags[]", tagID)
	}

	for _, tagID := range excludedTagsIDs {
		params.Add("excludedTags[]", tagID)
	}
	params.Add("limit", fmt.Sprintf("%d", limit))

	mangaReq := fmt.Sprintf("%s/manga?%s", c.BaseURL, params.Encode())
	fmt.Println(mangaReq)
	mangaResp, err := http.Get(mangaReq)
	if err != nil {
		return nil, err
	}
	defer mangaResp.Body.Close()

	var mangaData api.Manga
	mangaErr := json.NewDecoder(mangaResp.Body).Decode(&mangaData)
	if mangaErr != nil {
		return nil, err
	}

	var mangaList []api.MangaData
	for _, manga := range mangaData.Data {
		mangaList = append(mangaList, manga)
	}

	return mangaList, nil
}

func extractTagIds(tags api.TagsList, includedTagNames, excludedTagNames []string) ([]string, []string) {
	var includedTagIDs []string
	var excludedTagIDS []string

	for _, tag := range tags.Data {
		tagName, ok := tag.Attributes.Name["en"]

		if ok {
			if contains(includedTagNames, tagName) {
				includedTagIDs = append(includedTagIDs, tag.ID.String())
			} else if contains(excludedTagNames, tagName) {
				excludedTagIDS = append(excludedTagIDS, tag.ID.String())
			}
		}
	}
	fmt.Printf("in extract tag ids func: %v, %v\n", includedTagIDs, excludedTagIDS)

	return includedTagIDs, excludedTagIDS
}

// compare the names ignoring case sensitivity, should rename function to something more practical
func contains(tagsNamesArr []string, tagName string) bool {
	for _, v := range tagsNamesArr {
		if strings.EqualFold(v, tagName) {
			return true
		}
	}
	return false
}
