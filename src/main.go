package main

/* TODO:
   1.) learn why the relationships array struct comes out blank.
	- the relationships array is pointing to other manga that have a relation with the series you looked up
	- and it will show what type of manga it has a relation with. Negima doujinshi -> to Negima series etc
   2.) figure out how to separate these into fuctions
   3.) really learn these standard libraries
*/
import (
	"encoding/json"
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

func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{},
		BaseURL:    baseUrl,
	}
}

func (c *Client) SearchManga(title string, limit int) ([]api.MangaData, error) {
	endpoint := fmt.Sprintf("%s/manga", c.BaseURL)
	params := url.Values{}
	params.Add("title", title)
	params.Add("limit", fmt.Sprintf("%d", limit))
	fmt.Println(params)

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
		fmt.Println(item.Attributes.Title)
	}
	test, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(test))
	return mangaList, nil
}

func (c *Client) searchByTags(includedTags, excludedTags []string, limit int) ([]api.MangaData, error) {
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

func main() {
	client := NewClient()
	includedTags := []string{"comedy", "horror"}
	excludedTags := []string{"action"}
	manga, _ := client.searchByTags(includedTags, excludedTags, 3)
	for _, m := range manga {
		fmt.Println("in main: ", m.Attributes.Title)
	}
}
