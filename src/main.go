package main

import (
	"fmt"

	"github.com/saul178/manga-library-proj/src/tests"
)

func main() {
	client := tests.TestClient()
	//	includedTags := []string{"comedy", "horror"}
	//	excludedTags := []string{"action"}
	// client.GetMangaVolumesInfo("Heart no Kuni no Alice")
	// client.SearchByTags(includedTags, excludedTags, 4)
	manga, _ := client.SearchManga("heart no kuni no alice", 10)
	for _, m := range manga {
		fmt.Println("manga title: ", m.Attributes.Title)
	}
}
