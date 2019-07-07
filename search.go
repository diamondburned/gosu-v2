package osu

import (
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
	Mode     Mode           `schema:"m"`
	Category SearchCategory `schema:"s"`
}

const SearchEndpoint = "https://osu.ppy.sh/beatmapsets/search"

func SearchBeatmaps(opts SearchOpts) (*SearchBeatmapsResults, error) {
	var form url.Values

	if err := Schema.Encode(opts, form); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", SearchEndpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Form = form

	r, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	// todo
}
