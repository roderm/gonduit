package responses

import (
	"github.com/uber/gonduit/util"
)

// DifferentialRevisionSearchResponse contains fields that are in server
// response to differential.revision.search.
type DifferentialRevisionSearchResponse struct {
	// Data contains search results.
	Data []*DifferentialRevisionSearchResponseItem `json:"data"`

	// Curson contains paging data.
	Cursor SearchCursor `json:"cursor,omitempty"`
}

// DifferentialRevisionSearchResponseItem contains information about a
// particular search result.
type DifferentialRevisionSearchResponseItem struct {
	SearchResponse
	Fields      DifferentialRevisionSearchResponseItemFields `json:"fields"`
	Attachments DifferentialRevisionSearchAttachments        `json:"attachments"`
}

type DifferentialRevisionSearchResponseItemFields struct {
	Title          string                     `json:"title"`
	URI            string                     `json:"uri"`
	AuthorPHID     string                     `json:"authorPHID"`
	Status         DifferentialRevisionStatus `json:"status"`
	RepositoryPHID string                     `json:"repositoryPHID"`
	DiffPHID       string                     `json:"diffPHID"`
	Summary        string                     `json:"summary"`
	TestPlan       string                     `json:"testPlan"`
	IsDraft        bool                       `json:"isDraft"`
	HoldAsDraft    bool                       `json:"holdAsDraft"`
	DateCreated    util.UnixTimestamp         `json:"dateCreated"`
	DateModified   util.UnixTimestamp         `json:"dateModified"`
}

// DifferentialRevisionStatus represents item status returned by response.
type DifferentialRevisionStatus struct {
	Value  string `json:"value"`
	Name   string `json:"name"`
	Closed bool   `json:"closed"`
}

// DifferentialRevisionSearchAttachments holds possible attachments for the API
// method.
type DifferentialRevisionSearchAttachments struct {
	Reviewers   SearchAttachmentReviewers   `json:"reviewers"`
	Subscribers SearchAttachmentSubscribers `json:"subscribers"`
	Projects    SearchAttachmentProjects    `json:"projects"`
}
