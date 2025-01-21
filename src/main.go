package main

import (
	"encoding/json"
	"fmt"

	"github.com/saul178/manga-library-proj/src/tests"
)

func main() {
	client := tests.TestClient()
	//	includedTags := []string{"comedy", "horror"}
	//	excludedTags := []string{"action"}
	// client.GetMangaVolumesInfo("Heart no Kuni no Alice")
	// client.SearchByTags(includedTags, excludedTags, 4)
	manga, _ := client.SearchManga("Monochrome Days", 1)
	test, _ := json.MarshalIndent(manga, "", " ")
	fmt.Println(string(test))
}
