package requests

import "github.com/uber/gonduit/util"

// FileDownloadRequest represents a call to file.download.
type FileDownloadRequest struct {
	PHID string `json:"phid"`
	Request
}

// FileSearchRequest contains the parameters for file.serach (https://secure.phabricator.com/conduit/method/file.search/)
type FileSearchRequest struct {
	QueryKey    string                   `json:"queryKey"`
	Contraints  *QueryConstraintsRequest `json:"constraints"`
	Order       []string                 `json:"order"`
	Attachments *FileAttachmentRequest   `json:"attachments"`
	Request
}

// FileAttachmentRequest has additional fields to query (https://secure.phabricator.com/conduit/method/file.search/)
type FileAttachmentRequest struct {
	Subscribers bool `json:"subscribers"`
}

// QueryConstraintsRequest filtering for file.search (https://secure.phabricator.com/conduit/method/file.search/)
type QueryConstraintsRequest struct {
	IDs          []uint64            `json:"ids"`
	PHIDs        []string            `json:"phids"`
	AuthorPHIDs  []string            `json:"authorPHIDs"`
	Explicit     *bool               `json:"explicit"`
	CreatedStart *util.UnixTimestamp `json:"createdStart"`
	CreatedEnd   *util.UnixTimestamp `json:"createdEnd"`
	Name         string              `json:"name"`
	Subscribers  []string            `json:"subscribers"`
}
