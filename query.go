package stackexchange

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

type AdvancedSearch struct {
	Accepted  bool       `json:",omitempty"`
	Answers   int        `json:",omitempty"`
	Body      string     `json:",omitempty"`
	Closed    bool       `json:",omitempty"`
	Migrated  bool       `json:",omitempty"`
	Notice    bool       `json:",omitempty"`
	NotTagged []string   `json:",omitempty"`
	Tagged    []string   `json:",omitempty"`
	Title     string     `json:",omitempty"`
	User      string     `json:",omitempty"`
	URL       string     `json:",omitempty"`
	Views     int        `json:",omitempty"`
	Wiki      bool       `json:",omitempty"`
	SortBy    SortMethod `json:",omitempty"`
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

func NewClient(client *OAuth) *Client {
	return &Client{
		oauth: client,
	}
}
