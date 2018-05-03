package onedrive

import (
	"errors"
	"net/http"
	"time"
)

// Exists return true if the file exists and the download url
func (o Onedrive) Exists(path string) (bool, string, error) {
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	fileUrl := o.graphUrl + path
	req, err := http.NewRequest(http.MethodGet, fileUrl, nil)
	if err != nil {
		return false, "", err
	}

	req.Header.Set("Authorization", "bearer"+o.token)
	res, err := httpClient.Do(req)
	if err != nil {
		return false, "", err
	}

	if res.StatusCode != 302 {
		return false, "", errors.New("file not found or unauthenticated " + string(res.StatusCode))
	}

	return true, res.Header.Get("Location"), nil
}
