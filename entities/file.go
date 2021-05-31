package entities

import "github.com/uber/gonduit/util"

// File is a picture or document in phabricator
type File struct {
	Name         string                 `json:"name"`
	URI          string                 `json:"uri"`
	DataURI      string                 `json:"dataURI"`
	Size         uint64                 `json:"size"`
	DateCreated  util.UnixTimestamp     `json:"dateCreated"`
	DateModified util.UnixTimestamp     `json:"dateModified"`
	Policy       map[string]interface{} `json:"policy"`
}
