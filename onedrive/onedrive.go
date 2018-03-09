package onedrive

import "errors"

const (
	graphUrl   = "https://graph.microsoft.com"
	apiVersion = "1.0"
	entity     = "me"
	drivePath  = "root:"
)

type Onedrive struct {
	token    string
	graphUrl string
}

// NewOnedrive create new Onedrive client.
// token argument needs to be the retrieved bearer token from the graph api,
// for example: bearer <access_token>
//
func NewOnedrive(token string) (*Onedrive, error) {
	if token == "" {
		return nil, errors.New("token cannot be empty")
	}
	onedrive := &Onedrive{
		token:    token,
		graphUrl: constructOneDriveUrl(),
	}
	return onedrive, nil
}

func constructOneDriveUrl() string {
	return graphUrl + "/" + apiVersion + "/" + entity + "/" + drivePath
}
