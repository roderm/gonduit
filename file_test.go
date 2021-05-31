package gonduit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/gonduit/core"
	"github.com/uber/gonduit/entities"
	"github.com/uber/gonduit/requests"
	"github.com/uber/gonduit/responses"
	"github.com/uber/gonduit/test/server"
)

const fileResponseJSON = `{
	"result": {
		"data": [
			{
				"id": 1000,
				"type": "FILE",
				"phid": "PHID-FILE-abcdefg1234",
				"fields": {
					"name": "logo.png",
					"uri": "https:\/\/www.example.com\/F1000",
					"dataURI": "https:\/\/www.example.com\/file\/data\/3aqt62dmshe2nri2pmeg\/PHID-FILE-qbxcb7rpqfokykht3mdh\/image.png",
					"size": 16392
				}
			}
		]
	}
}`

func TestFileSearch(t *testing.T) {
	s := server.New()
	defer s.Close()
	s.RegisterCapabilities()
	response := server.ResponseFromJSON(fileResponseJSON)
	s.RegisterMethod(FileSearchMethod, http.StatusOK, response)

	c, err := Dial(s.GetURL(), &core.ClientOptions{
		APIToken: "some-token",
	})
	assert.Nil(t, err)
	req := requests.FileSearchRequest{}
	resp, err := c.FileSearch(req)
	assert.NoError(t, err)
	want := responses.FileSearchResponse{
		Data: []responses.FileResponse{
			{
				ID:   1000,
				Type: "FILE",
				PHID: "PHID-FILE-abcdefg1234",
				Fields: entities.File{
					Name:    "logo.png",
					URI:     "https://www.example.com/F1000",
					DataURI: "https://www.example.com/file/data/3aqt62dmshe2nri2pmeg/PHID-FILE-qbxcb7rpqfokykht3mdh/image.png",
					Size:    16392,
				},
			},
		},
	}
	assert.Equal(t, &want, resp)
}
