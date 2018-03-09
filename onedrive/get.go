package onedrive

import "errors"

// Get return a given file from Onedrive as byte array
// if no path is given it return an error
func (o Onedrive) Get(path string) ([]byte, error) {
	if path == "" {
		return nil, errors.New("path cannot be empty")
	}

	var content []byte

	return content, nil
}
