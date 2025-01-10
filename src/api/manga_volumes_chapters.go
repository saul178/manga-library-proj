// this is to get the a mangas series volumes and chapters data
package api

import "github.com/google/uuid"

type MangaVolumeResponse struct {
	Result  string                            `json:"result"`
	Volumes map[string]MangaVolumesProperties `json:"volumes"`
}

type MangaVolumesProperties struct {
	Volume   string                       `json:"volume"`
	Count    int                          `json:"count"`
	Chapters map[string]ChapterProperties `json:"chapters"`
}

type ChapterProperties struct {
	Chapter string      `json:"chapter"`
	ID      uuid.UUID   `json:"id"`
	Others  []uuid.UUID `json:"others"`
	Count   int         `json:"count"`
}
