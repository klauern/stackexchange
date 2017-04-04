package stackexchange

import "github.com/dghubble/sling"

type ResponseWrapper struct {
	Items          []interface{} `json:"items"`
	BackOffSeconds int           `json:"backoff,omitempty"`
	ErrorID        int           `json:"error_id,omitempty"`
	ErrorMessage   string        `json:"error_message,omitempty"`
	ErrorName      string        `json:"error_name,omitempty"`
	HasMore        bool          `json:"has_more"`
	Page           int           `json:"page,omitempty"`
	PageSize       int           `json:"page_size,omitempty"`
	QuotaMax       int           `json:"quota_max"`
	QuotaRemaining int           `json:"quota_remaining"`
	Total          int           `json:"total,omitempty"`
	Type           string        `json:"type,omitempty"`
}

type SortMethod int

const (
	Activity SortMethod = iota
	Creation
	Votes
	Relevance
)

func (m *SortMethod) String() string {
	switch m {
	case Activity:
		return "last_activity_date"
	case Creation:
		return "creation_date"
	case Votes:
		return "score"
	case Relevance:
		return "relevance"
	default:
		return "last_activity_date"
	}
}

type AdvancedSearch struct {
	Accepted  bool       `url:"accepted,omitempty"`
	Answers   int        `url:"answers,omitempty"`
	Body      string     `url:"body,omitempty"`
	Closed    bool       `url:"closed,omitempty"`
	Migrated  bool       `url:"migrated,omitempty"`
	Notice    bool       `url:"notice,omitempty"`
	NotTagged []string   `url:"nottagged,omitempty"`
	Tagged    []string   `url:"tagged,omitempty"`
	Title     string     `url:"title,omitempty"`
	User      string     `url:"user,omitempty"`
	URL       string     `url:"url,omitempty"`
	Views     int        `url:"views,omitempty"`
	Wiki      bool       `url:"wiki,omitempty"`
	SortBy    SortMethod `url:"sort,omitempty"`
}

const redirectURI = "https://stackexchange.com/oauth/login_success"
const accessToken = "Sh4G4lfjrmsgFy7RySnrUQ))"
const expires = 86399

var BaseURI = "https://api.stackexchange.com/2.2/"

type OAuth struct {
	clientID     int    `json:"client_id"`
	clientSecret string `json:"client_secret"`
	key          string `json:"key"`
}

type Client struct {
	oauth *OAuth
}

func GetThing() {
	slinger := sling.New()
	slinger.Get("/")
}

func NewClient(client *OAuth) *Client {
	return &Client{
		oauth: client,
	}
}
