package osu

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func authenticate(t *testing.T) *Session {
	file, err := ioutil.ReadFile("_testdata/credentials.json")
	if err != nil {
		t.Fatal(err)
	}

	var creds = struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := json.Unmarshal(file, &creds); err != nil {
		t.Fatal(err)
	}

	s, err := NewSession(creds.Username, creds.Password)
	if err != nil {
		t.Fatal(err)
	}

	if err := s.Login(); err != nil {
		t.Fatal(err)
	}

	return s
}

func TestSession(t *testing.T) {
	s := authenticate(t)
	spew.Dump(s)
}

func TestDownloadBeatmapset(t *testing.T) {
	s := authenticate(t)

	// https://osu.ppy.sh/beatmapsets/931596
	body, err := s.DownloadBeatmap("931596")
	if err != nil {
		t.Fatal(err)
	}

	defer body.Close()

	file, err := os.OpenFile(
		"_testdata/test.osz",
		os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, os.ModePerm,
	)

	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	if _, err := io.Copy(file, body); err != nil {
		t.Fatal(err)
	}
}
