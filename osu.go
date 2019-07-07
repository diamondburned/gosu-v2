package osu

import (
	"net/http"
	"time"

	"github.com/gorilla/schema"
)

var (
	Schema = schema.NewEncoder()
	Client = http.Client{
		Timeout: 5 * time.Second,
	}
)
