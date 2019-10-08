package requests

// DifferentialGetCommitPathsRequest represents a request to the
// differential.getcommitpaths call.
type DifferentialGetCommitPathsRequest struct {
	RevisionID uint64 `json:"revision_id"`
	Request
}
