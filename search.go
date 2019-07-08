package osu

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// SearchCategory is used for searching map statuses
// Example: "Raned & Approved", "Loved", etc
type SearchCategory string

const (
	SearchCategoryRanked        SearchCategory = "ranked"
	SearchCategoryFavorites     SearchCategory = "favorites"
	SearchCategoryQualified     SearchCategory = "qualified"
	SearchCategoryPendingAndWIP SearchCategory = "pending"
	SearchCategoryGraveyard     SearchCategory = "graveyard"
	SearchCategoryMyMaps        SearchCategory = "mine"
	SearchCategoryAny           SearchCategory = "any"
	SearchCategoryLoved         SearchCategory = "loved"
)

type SearchOpts struct {
	Query    string         `schema:"q,required"`
	Mode     ModeInt        `schema:"m"`
	Category SearchCategory `schema:"s"`
}

const SearchEndpoint = "https://osu.ppy.sh/beatmapsets/search"

type SearchBeatmapsResults struct {
	Beatmapsets           []Beatmapset `json:"beatmapsets"`
	Cursor                *Cursor      `json:"cursor,omitempty"`
	RecommendedDifficulty float64      `json:"recommended_difficulty"`
	Total                 int64        `json:"total"`
}

func (s *Session) SearchBeatmaps(opts SearchOpts) (*SearchBeatmapsResults, error) {
	if opts.Category == "" {
		opts.Category = SearchCategoryRanked
	}

	var form = url.Values{}

	if err := Schema.Encode(opts, form); err != nil {
		return nil, err
	}

	println(form.Encode())

	req, err := http.NewRequest(
		"GET", SearchEndpoint+"?"+form.Encode(), nil,
	)

	if err != nil {
		return nil, err
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	var bmr SearchBeatmapsResults

	if err := json.NewDecoder(res.Body).Decode(&bmr); err != nil {
		return nil, err
	}

	return &bmr, nil
}
