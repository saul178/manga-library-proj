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

type Manga struct {
	ID    string            `json:"id"`
	Title map[string]string `json:"title"`
}

func (c *Client) SearchManga(title string, limit int) ([]Manga, error) {
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

	var result struct {
		Data []struct {
			ID         string `json:"id"`
			Attributes struct {
				Title map[string]string `json:"title"`
			} `json:"attributes"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	var mangaList []Manga

	for _, item := range result.Data {
		mangaList = append(mangaList, Manga{
			ID:    item.ID,
			Title: item.Attributes.Title,
		})
	}
	return mangaList, nil
}

func main() {
	client := NewClient()
	manga, err := client.SearchManga("negima", 5)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for _, m := range manga {
		fmt.Printf("ID: %s, Title: %v\n", m.ID, m.Title["en"])
	}
}
