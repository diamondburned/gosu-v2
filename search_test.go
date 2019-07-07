package osu

import (
	"testing"
)

func TestSearchBeatmaps(t *testing.T) {
	r, err := SearchBeatmaps(SearchOpts{
		Query: "Himeringo Sotarks",
	})

	if err != nil {
		t.Error(err)
	}

	if r.Beatmapsets[0].ID != 903106 {
		t.Errorf("First beatmap is wrong, expecting 903106, got %v", r.Beatmapsets[0].ID)
	}
}
