package osu

import (
	"errors"
	"net/url"
)

type loginForm struct {
	Username string `schema:"username"`
	Password string `schema:"password"`

	XSRFToken string `schema:"_token"`
}

// Login tries to authenticate with the osu! servers and set the
// tokens in place
func (s *Session) Login() error {
	URL, _ := url.Parse("https://osu.ppy.sh/home")

	if s.login.XSRFToken == "" {
		// Access this site for the token
		r, err := s.get(URL.String())
		if err != nil {
			return err
		}

		r.Body.Close()

		for _, c := range s.client.Jar.Cookies(URL) {
			if c.Name == "XSRF-TOKEN" {
				s.login.XSRFToken = c.Value
				break
			}
		}
	}

	if s.login.XSRFToken == "" {
		return errors.New("Failed to get XSRF token")
	}

	var v = url.Values{}
	if err := s.schema.Encode(s, v); err != nil {
		return err
	}

	r, err := s.post("https://osu.ppy.sh/session?"+v.Encode(), nil)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	if r.StatusCode < 200 || r.StatusCode > 299 {
		return NewResponseError(r, nil)
	}

	// We do nothing at this point, because the request above should've set the
	// cookiejar to have the token.

	for _, c := range s.client.Jar.Cookies(URL) {
		if c.Name == "osu_session" {
			s.SessionToken = c.Value
			break
		}
	}

	return nil
}
