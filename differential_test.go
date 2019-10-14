package gonduit

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/uber/gonduit/core"
	"github.com/uber/gonduit/requests"
	"github.com/uber/gonduit/responses"
	"github.com/uber/gonduit/test/server"
)

func TestDifferentialGetCommitPaths(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()

	s.RegisterMethod(DifferentialGetCommitPathsMethod, http.StatusOK, gin.H{
		"result": []string{
			"differential.go",
			"differential_test.go",
		},
	})

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.DifferentialGetCommitPathsRequest{RevisionID: 123}
	resp, err := c.DifferentialGetCommitPaths(req)
	assert.NoError(t, err)
	want := &responses.DifferentialGetCommitPathsResponse{
		"differential.go",
		"differential_test.go",
	}
	assert.Equal(t, want, resp)
}

func TestDifferentialGetCommitMessage(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()

	s.RegisterMethod(DifferentialGetCommitMessageMethod, http.StatusOK, gin.H{
		"result": "Commit description.",
	})

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.DifferentialGetCommitMessageRequest{
		RevisionID: 123,
	}
	resp, err := c.DifferentialGetCommitMessage(req)
	assert.NoError(t, err)
	want := responses.DifferentialGetCommitMessageResponse("Commit description.")
	assert.Equal(t, &want, resp)
}
