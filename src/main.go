package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/saul178/manga-library-proj/src/helpers"
)

// TODO: learn why the titles come out blank, figure out how to separate these into fuctions and
// really learn these standard libraries

func main() {
	title := "Kanojyo to Himitsu to Koimoyou"
	baseUrl := "https://api.mangadex.org"

	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal("you did something wrong", err)
	}

	u.Path += "/manga"

	parameters := url.Values{}
	parameters.Add("title", title)
	u.RawQuery = parameters.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal("you're actually bad at this", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var manga helpers.MangaData
	json.Unmarshal(body, &manga)

	for index, m := range manga.Data {
		fmt.Println(index, " ", m.ID)
	}
}
