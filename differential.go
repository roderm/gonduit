package gonduit

import (
	"github.com/uber/gonduit/requests"
	"github.com/uber/gonduit/responses"
)

// DifferentialQuery performs a call to differential.query.
func (c *Conn) DifferentialQuery(
	req requests.DifferentialQueryRequest,
) (*responses.DifferentialQueryResponse, error) {
	var res responses.DifferentialQueryResponse

	if err := c.Call("differential.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DifferentialQueryDiffs performs a call to differential.querydiffs.
func (c *Conn) DifferentialQueryDiffs(
	req requests.DifferentialQueryDiffsRequest,
) (*responses.DifferentialQueryDiffsResponse, error) {
	var res responses.DifferentialQueryDiffsResponse

	if err := c.Call("differential.querydiffs", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// DifferentialGetCommitPathsMethod is method name on Phabricator API.
const DifferentialGetCommitPathsMethod = "differential.getcommitpaths"

// DifferentialGetCommitPaths performs a call to differential.getcommitpaths.
func (c *Conn) DifferentialGetCommitPaths(
	req requests.DifferentialGetCommitPathsRequest,
) (*responses.DifferentialGetCommitPathsResponse, error) {
	var res responses.DifferentialGetCommitPathsResponse

	if err := c.Call(DifferentialGetCommitPathsMethod, &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
