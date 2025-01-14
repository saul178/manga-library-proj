package main

import (
	"fmt"

	"github.com/saul178/manga-library-proj/src/tests"
)

func main() {
	client := tests.TestClient()
	//	includedTags := []string{"comedy", "horror"}
	//	excludedTags := []string{"action"}
	client.GetMangaVolumesInfo("Maria the Virgin Witch")
	manga, _ := client.SearchManga("Maria the Virgin Witch", 10)
	for _, m := range manga {
		fmt.Println("manga title: ", m.Attributes.Title)
	}
}
