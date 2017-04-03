package stackexchange

type ResponseWrapper struct {
	Items []struct {
		IncludedFields []string `json:"included_fields"`
		FilterType     string   `json:"filter_type"`
		Filter         string   `json:"filter"`
	} `json:"items"`
	BackOffSeconds int    `json:"backoff,omitempty"`
	ErrorID        int    `json:"error_id,omitempty"`
	ErrorMessage   string `json:"error_message,omitempty"`
	ErrorName      string `json:"error_name,omitempty"`
	HasMore        bool   `json:"has_more"`
	Page           int    `json:"page,omitempty"`
	PageSize       int    `json:"page_size,omitempty"`
	QuotaMax       int    `json:"quota_max"`
	QuotaRemaining int    `json:"quota_remaining"`
	Total          int    `json:"total,omitempty"`
	Type           string `json:"type,omitempty"`
}
