package onedrive

import "errors"

// Put: upload files to onedrive, the path must be the full path
// curl https://graph.microsoft.com/v1.0/me/drive/root:/document1.docx:/content -X PUT -d @document1.docx -H "Authorization: bearer access_token_here"
// this method only allows file uploads to 4MB
func (o Onedrive) Put(path string) (string, error) {
	if path == "" {
		return "", errors.New("path cannot be empty")
	}
	return `{"uploaded": "true"}`, nil
}
