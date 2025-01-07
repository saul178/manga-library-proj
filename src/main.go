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

	var result api.MangaList

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

	// params := url.Values{}
	// params.Add("includedTags[]", includedTags[]) // for this to work i need the id of the tags so that they map to related manga
}

func extractTagIds(includedTagNames, excludedTagNames []string) ([]string, []string) {
	var tags api.TagsList
	var includedTagIDs []string
	var excludedTagIDS []string

	for _, tag := range tags.Data {
		name, exists := tag.Attributes.Name["en"]
		if exists {
		}
	}

	return includedTagIDs, excludedTagIDS
}

func main() {
	client := NewClient()
	manga, err := client.SearchManga("negima", 1)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for _, m := range manga {
		fmt.Println("in main: ", m)
	}
}
