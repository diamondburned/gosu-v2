package osu

type SearchBeatmapsResults struct {
	Beatmapsets           []Beatmapset `json:"beatmapsets"`
	RecommendedDifficulty float64      `json:"recommended_difficulty"`
	Total                 int64        `json:"total"`
}

type Beatmapset struct {
	ID                int64     `json:"id"`
	Title             string    `json:"title"`
	Artist            string    `json:"artist"`
	PlayCount         int64     `json:"play_count"`
	FavouriteCount    int64     `json:"favourite_count"`
	HasFavourited     bool      `json:"has_favourited"`
	SubmittedDate     string    `json:"submitted_date"`
	LastUpdated       string    `json:"last_updated"`
	RankedDate        string    `json:"ranked_date"`
	Creator           string    `json:"creator"`
	UserID            int64     `json:"user_id"`
	BPM               float64   `json:"bpm"`
	Source            string    `json:"source"`
	Covers            Covers    `json:"covers"`
	PreviewURL        string    `json:"preview_url"`
	Tags              string    `json:"tags"`
	Video             bool      `json:"video"`
	Storyboard        bool      `json:"storyboard"`
	Ranked            int64     `json:"ranked"`
	Status            Status    `json:"status"`
	HasScores         bool      `json:"has_scores"`
	DiscussionEnabled bool      `json:"discussion_enabled"`
	CanBeHyped        bool      `json:"can_be_hyped"`
	Hype              Hype      `json:"hype"`
	Nominations       Hype      `json:"nominations"`
	LegacyThreadURL   string    `json:"legacy_thread_url"`
	Beatmaps          []Beatmap `json:"beatmaps"`
}

type Beatmap struct {
	ID               int64       `json:"id"`
	BeatmapsetID     int64       `json:"beatmapset_id"`
	Mode             Mode        `json:"mode"`
	ModeInt          int64       `json:"mode_int"`
	Convert          interface{} `json:"convert"`
	DifficultyRating float64     `json:"difficulty_rating"`
	Version          string      `json:"version"`
	TotalLength      int64       `json:"total_length"`
	CS               float64     `json:"cs"`
	Drain            float64     `json:"drain"`
	Accuracy         float64     `json:"accuracy"`
	Ar               float64     `json:"ar"`
	Playcount        int64       `json:"playcount"`
	Passcount        int64       `json:"passcount"`
	CountCircles     int64       `json:"count_circles"`
	CountSliders     int64       `json:"count_sliders"`
	CountSpinners    int64       `json:"count_spinners"`
	CountTotal       int64       `json:"count_total"`
	LastUpdated      string      `json:"last_updated"`
	Ranked           int64       `json:"ranked"`
	Status           Status      `json:"status"`
	URL              string      `json:"url"`
	DeletedAt        interface{} `json:"deleted_at"`
}

type Covers struct {
	Cover       string `json:"cover"`
	Cover2X     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2X      string `json:"card@2x"`
	List        string `json:"list"`
	List2X      string `json:"list@2x"`
	Slimcover   string `json:"slimcover"`
	Slimcover2X string `json:"slimcover@2x"`
}

type Hype struct {
	Current  int64 `json:"current"`
	Required int64 `json:"required"`
}

// Mode contains the osu! gamemode in strings
type Mode string

const (
	ModeStandard Mode = "osu"
	ModeTaiko    Mode = "taiko"
	ModeMania    Mode = "mania"
	ModeCatch    Mode = "fruits"
)

type ModeInt int

const (
	IntStandard ModeInt = iota
	IntTaiko
	IntCatch
	IntMania
)

/*
var ModeInts = map[Mode]ModeInt{
	ModeStandard: IntStandard,
	ModeTaiko:    IntTaiko,
	ModeCatch:    IntCatch,
	ModeMania:    IntMania,
}
*/

// Status is the map status (ranked, pending, etc)
type Status string

const (
	StatusRanked    Status = "ranked"
	StatusQualified Status = "qualified"
	StatusLoved     Status = "loved"
	StatusWIP       Status = "wip"
	StatusPending   Status = "pending"
	StatusGraveyard Status = "graveyard"
)
