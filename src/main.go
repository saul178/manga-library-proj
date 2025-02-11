package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/saul178/manga-library-proj/src/api"
)

func main() {
	client := api.DexClient()
	//	includedTags := []string{"comedy", "horror"}
	//	excludedTags := []string{"action"}
	// client.GetMangaVolumesInfo("Heart no Kuni no Alice")
	// client.SearchByTags(includedTags, excludedTags, 4)

	id, _ := uuid.Parse("a77742b1-befd-49a4-bff5-1ad4e6b0ef7b")
	include := []string{"manga", "cover_art", "author"}
	// title := "chainsaw man"
	// limit := 1
	params := url.Values{}
	// params.Add("title", title)
	// params.Add("limit", fmt.Sprintf("%d", limit))
	for i := 0; i < len(include); i++ {
		params.Add("includes[]", include[i])
	}
	ctx := context.Background()

	manga, _ := client.GetManga(ctx, id, params)
	test, _ := json.MarshalIndent(manga, "", " ")
	fmt.Println(string(test))
}
