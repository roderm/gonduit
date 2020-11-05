package responses

// SearchResponse struct holds fields which are common for all *.search API
// methods.
type SearchResponse struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	PHID string `json:"phid"`
	SearchCursor
}

// SearchCursor holds paging information on responses from *.search API methods.
type SearchCursor struct {
	Limit  uint64 `json:"limit"`
	After  string `json:"after"`
	Before string `json:"before"`
}
