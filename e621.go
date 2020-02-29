package e926

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type JsonTime struct {
	JsonClass string
	S         int
	N         int
}

type Post struct {
	ID            int      `json:"id"`
	Tags          string   `json:"tags"`
	LockedTags    string   `json:"locked_tags"`
	Description   string   `json:"description"`
	CreatedAt     JsonTime `json:"created_at"`
	CreatorID     int      `json:"creator_id"`
	Author        string   `json:"author"`
	Change        int      `json:"change"`
	Source        string   `json:"source"`
	Score         int      `json:"score"`
	FavCount      int      `json:"fav_count"`
	Md5           string   `json:"md5"`
	FileSize      int      `json:"file_size"`
	FileURL       string   `json:"file_url"`
	FileExt       string   `json:"file_ext"`
	PreviewURL    string   `json:"preview_url"`
	PreviewWidth  int      `json:"preview_width"`
	PreviewHeight int      `json:"preview_height"`
	SampleURL     string   `json:"sample_url"`
	SampleWidth   int      `json:"sample_width"`
	SampleHeight  int      `json:"sample_height"`
	Rating        string   `json:"rating"`
	Status        string   `json:"status"`
	Width         int      `json:"width"`
	Height        int      `json:"height"`
	HasComments   bool     `json:"has_comments"`
	HasNotes      bool     `json:"has_notes"`
	HasChildren   bool     `json:"has_children"`
	Children      string   `json:"children"`
	ParentID      int      `json:"parent_id"`
	Artist        []string `json:"artist"`
	Sources       []string `json:"sources"`
}

type Session struct {
	BaseURL   string
	Username  string
	UserAgent string
	ApiKey    string
	Client    *http.Client
	waiting bool
	waitTicker *time.Ticker
}

// NewSession
func NewSession(domain string, userAgent string) *Session {
	return &Session{
		UserAgent: userAgent,
		Client: &http.Client{},
		BaseURL: "https://"+domain,
	}
}

func (s *Session) WaitBetweenRequests() {
	s.waiting = true
	s.waitTicker = time.NewTicker(time.Millisecond*500)
}

func (s *Session) PostURL(p *Post) string {
	return fmt.Sprintf(s.BaseURL+"/post/show/%v", p.ID)
}

func (s *Session) Get(url string, params map[string]string) (*http.Response, error) {
	if s.waiting {
		<- s.waitTicker.C
	}
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

func (s *Session) GetPosts(tags []string, limit int) (posts []*E621Post) {
	resp, err := s.Get("/post/index.json", map[string]string{
		"tags":  strings.Join(tags, " "),
		"limit": strconv.Itoa(limit),
	})

	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	json.Unmarshal(data, &posts)
	return posts

}
