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
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/saul178/manga-library-proj/src/api"
)

const baseUrl = "https://api.mangadex.org"

func searchByTag(includedTags []string, excludedTags []string) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal("you did something wrong ", err)
	}

	u.Path += "/manga/tag"
}

func main() {
	title := "negima"

	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal("you did something wrong ", err)
	}

	u.Path += "/manga"

	parameters := url.Values{}
	parameters.Add("title", title)
	u.RawQuery = parameters.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal("you're actually bad at this ", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var manga api.MangaData
	json.Unmarshal(body, &manga)

	indentedJson, err := json.MarshalIndent(manga, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(indentedJson))

	// holdMangaData := manga.Data
	//
	//	for index, m := range holdMangaData {
	//		fmt.Println(index, " ", m.Attributes.Title["en"], m.Attributes.Links["mal"])
	//		// fmt.Println(m.Attributes.Description)
	//	}
	//
	//	for _, j := range holdMangaData {
	//		for _, alt := range j.Attributes.AltTitles {
	//			title, okStatus := alt["ja"]
	//			if okStatus && title != "" {
	//				fmt.Println(title)
	//			}
	//		}
	//	}
}
