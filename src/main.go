package main

import (
	// "fmt"

	"github.com/saul178/manga-library-proj/src/tests"
)

func main() {
	client := tests.TestClient()
	//	includedTags := []string{"comedy", "horror"}
	//	excludedTags := []string{"action"}
	client.GetMangaVolumesInfo("chainsaw man")
	//	for _, m := range manga {
	//		fmt.Println("in main: ", m.ID)
	//	}
}
