package requests

import "github.com/uber/gonduit/constants"

// DifferentialGetCommitMessageRequest represents a request to the
// differential.getcommitmessage call.
type DifferentialGetCommitMessageRequest struct {
	RevisionID uint64                                         `json:"revision_id"`
	Fields     []string                                       `json:"fields"`
	Edit       constants.DifferentialGetCommitMessageEditType `json:"edit"`
	Request
}
