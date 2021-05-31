package gonduit

import (
	"github.com/uber/gonduit/requests"
	"github.com/uber/gonduit/responses"
)

// FileDownload performs a call to file.download.
func (c *Conn) FileDownload(
	req requests.FileDownloadRequest,
) (*responses.FileDownloadResponse, error) {
	var res responses.FileDownloadResponse

	if err := c.Call("file.download", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// FileSearchMethod is the method name on API.
const FileSearchMethod = "file.search"

// FileSearch performs a call to file.search.
func (c *Conn) FileSearch(
	req requests.FileSearchRequest,
) (*responses.FileSearchResponse, error) {
	var res responses.FileSearchResponse

	if err := c.Call(FileSearchMethod, &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
