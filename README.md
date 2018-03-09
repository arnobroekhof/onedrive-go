# OneDrive Go Client
## NOTE: still work in progress 

[![Build Status](https://travis-ci.org/arnobroekhof/onedrive-go.svg?branch=master)](https://travis-ci.org/arnobroekhof/onedrive-go)

### Supported methods

* Get file
* Put file


## Rest Calls example ( cURL )
### Upload
curl https://graph.microsoft.com/v1.0/me/drive/root:/document1.docx:/content -X PUT -d @document1.docx -H "Authorization: bearer access_token_here"
###  Download
curl https://graph.microsoft.com/v1.0/me/drive/root:/document1.docx:/content -X GET  -H "Authorization: bearer access_token_here"
### Convert to PDF
curl https://graph.microsoft.com/v1.0/me/drive/root:/document1.docx:/content?format=pdf -o document1.pdf -H "Authorization: bearer access_token_here"
