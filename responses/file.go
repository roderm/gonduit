package responses

import "github.com/uber/gonduit/entities"

// FileDownloadResponse represents a response from calling file.download.
type FileDownloadResponse struct {
	Result string `json:"result"`
}

// FileResponse contains the basic of a file
type FileResponse struct {
	ID     int64         `json:"id"`
	Type   string        `json:"type"`
	PHID   string        `json:"phid"`
	Fields entities.File `json:"fields"`
}

// FileSearchResponse is the requests result
type FileSearchResponse struct {
	Data []FileResponse `json:"data"`
}

// GetData returns the array of files
func (r *FileSearchResponse) GetData() []FileResponse {
	return r.Data
}
