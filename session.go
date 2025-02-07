package osu

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/gorilla/schema"
)

// Session contains information about the osu! session
type Session struct {
	SessionToken string

	login  loginForm
	client *http.Client
	schema *schema.Encoder
}

// NewSession initializes a Session struct with an in-memory cookiejar
func NewSession(username, password string) (*Session, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	client.Jar = jar

	return &Session{
		login: loginForm{
			Username: username,
			Password: password,
		},
		client: client,
		schema: schema.NewEncoder(),
	}, nil
}

// SetCookiejar sets the cookiejar
// This should be ran before Login
func (s *Session) SetCookiejar(c http.CookieJar) {
	s.client.Jar = c
}

// SetSessionToken should only be used if you know what this is
// If the word "XSRF-TOKEN" doesn't ring a bell, do NOT touch this
func (s *Session) SetSessionToken(t string) {
	url, _ := url.Parse("https://osu.ppy.sh/home")

	cookies := s.client.Jar.Cookies(url)
	defer s.client.Jar.SetCookies(url, cookies)

	for _, c := range cookies {
		if c.Name == "osu_session" {
			c.Value = t
			return
		}
	}

	cookies = append(
		cookies,
		&http.Cookie{
			Name:  "osu_session",
			Value: t,
		},
	)
}

// DownloadBeatmap returns a body which is the content of the beatmap.
// You'll need to manually close the returned body
func (s *Session) DownloadBeatmap(beatmapsetID string) (io.ReadCloser, error) {
	r, err := s.get(
		"https://osu.ppy.sh/beatmapsets/" + beatmapsetID + "/download",
	)

	if err != nil {
		return nil, err
	}

	if r.StatusCode < 200 || r.StatusCode > 300 {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(
			"Invalid status code %d\n%s",
			r.StatusCode, string(body),
		)
	}

	return r.Body, nil
}

func (s *Session) get(url string) (*http.Response, error) {
	return s.client.Get(url)
}

func (s *Session) post(url string, body io.Reader) (*http.Response, error) {
	r, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	return s.client.Do(r)
}
