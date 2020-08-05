package e621

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const MAX_LIMIT = 320

type Session struct {
	BaseURL   string
	Username  string
	UserAgent string
	ApiKey    string
	Client    *http.Client
}

// NewSession
func NewSession(domain string, userAgent string) *Session {
	return &Session{
		UserAgent: userAgent,
		Client:    &http.Client{},
		BaseURL:   "https://" + domain,
	}
}
func (s *Session) PostUrl(p *Post) string {
	return s.BaseURL+"/posts/" + strconv.Itoa(p.ID)
}
func (s *Session) Get(url string, params map[string]string) (*http.Response, error) {
	params["password_hash"] = s.ApiKey
	params["login"] = s.Username

	req, err := http.NewRequest("GET", s.BaseURL+url, nil)

	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Add("User-Agent", s.UserAgent)

	return s.Client.Do(req)
}

func (s *Session) GetPosts(tags string, limit int) (posts PostsResponse, err error) {
	if limit > MAX_LIMIT {
		panic("Limit must be bellow " + strconv.Itoa(MAX_LIMIT))
	}

	resp, err := s.Get("/posts.json", map[string]string{
		"tags":  tags,
		"limit": strconv.Itoa(limit),
	})

	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &posts)
	return
}

func (s *Session) FindAliases(name string) (aliases []TagAlias, err error) {
	resp, err := s.Get("/tag_aliases.json", map[string]string{
		"search[name_matches]": name,
	})
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &aliases)
	return
}

func (s *Session) FindTag(name string) (tags []Tag, err error) {
	resp, err := s.Get("/tags.json", map[string]string{
		"search[name_matches]": name,
	})
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &tags)
	return
}
