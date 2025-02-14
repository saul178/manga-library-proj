// this is to make api calls to get mangas by searching based on tags etc.
package api

import (
	"context"
	"net/url"

	"github.com/google/uuid"
)

const (
	// this is to get the tags information such as the ID
	tagEndpoint = "/manga/tag"
	// hit this endpoint with the tag paramaters
	mangaEndpoint = "/manga"
)

type TagResponse struct {
	Result   string `json:"result"`
	Response string `json:"response"`
}

type TagsList struct {
	TagResponse
	Data   []TagsData `json:"data"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Total  int        `json:"total"`
}

type TagsData struct {
	ID            uuid.UUID          `json:"id"`
	Type          string             `json:"type"`
	Attributes    TagAttributes      `json:"attributes"`
	Relationships []TagRelationships `json:"relationships"`
}

type TagAttributes struct {
	Name        map[string]string `json:"name"`
	Description map[string]string `json:"description"`
	Group       string            `json:"group"`
	Version     int               `json:"version"`
}

type TagRelationships []struct {
	ID         uuid.UUID   `json:"id"`
	Type       string      `json:"type"`
	Related    string      `json:"related,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

/*
TODO: the whole point of the tags is that i should be able to look up manga by inputting specific tags and getting a list
of mangas in conjuction of the tags that were inputted. first need to hit the tag endpoint with our parameters for included
tags and excluded tags, grab the UUID's of these tags, then we send the request with the parameters to /manga

look in local stuff first, if not there, then request from mangadex.
*/

/*
NOTE: can filter tags by giving an array of including[] tags, and giving an array of excluding[] tags and the language
code, im essential writing an advance search bar for my app.
*/
func getTagIDs(ctx context.Context, includedTagNames, excludedTagNames []string) (map[string]uuid.UUID, error) {
}

func (s *MangadexService) SearchMangaByTags(ctx context.Context, params url.Values, langCode string) (*TagResponse, error) {
}
