// this is to get the a mangas series volumes and chapters data
package api

import "github.com/google/uuid"

type MangaVolumeResponse struct {
	Result  string         `json:"result"`
	Volumes []MangaVolumes `json:"volumes"`
}

type MangaVolumes struct {
	Properties []MangaVolumesProperties `json:"properties"`
}

type MangaVolumesProperties struct {
	Volume   string                `json:"volume"`
	Count    int                   `json:"count"`
	Chapters []MangaVolumeChapters `json:"chapters"`
}

type MangaVolumeChapters struct {
	Property []ChapterProperties `json:"property"`
	Count    int                 `json:"count"`
}

type ChapterProperties struct {
	Chapter string      `json:"chapter"`
	ID      uuid.UUID   `json:"id"`
	Others  []uuid.UUID `json:"others"`
	Count   int         `json:"count"`
}
