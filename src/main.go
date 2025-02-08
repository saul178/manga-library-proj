package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/saul178/manga-library-proj/src/api"
)

func main() {
	client := api.DexClient()
	//	includedTags := []string{"comedy", "horror"}
	//	excludedTags := []string{"action"}
	// client.GetMangaVolumesInfo("Heart no Kuni no Alice")
	// client.SearchByTags(includedTags, excludedTags, 4)

	title := "chainsaw man"
	limit := 1
	params := url.Values{}
	params.Add("title", title)
	ctx := context.Background()
	params.Add("limit", fmt.Sprintf("%d", limit))
	manga, _ := client.SearchMangas(ctx, params)
	test, _ := json.MarshalIndent(manga, "", " ")
	fmt.Println(string(test))
}
