# OneDrive Go Client
## NOTE: still work in progress 

[![Build Status](https://travis-ci.org/arnobroekhof/onedrive-go.svg?branch=master)](https://travis-ci.org/arnobroekhof/onedrive-go)

### Supported methods

* Get file
* Put file

### Middleware libraries

* Azure JWT authentication and verification

example:


create the following go code:

```go
package main

import (
	"github.com/arnobroekhof/onedrive-go/auth"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/test", auth.AzureJWTAuthMiddleware(TestHandler())).Methods(http.MethodGet)
	log.Println("Listening on http://127.0.0.1:8080")
	log.Fatalln(http.ListenAndServe("127.0.0.1:8080", r))

}

func TestHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(`{"message": "it works"}`))
	}
}
```

create request using cURL:

```bash
curl -v -XGET http://localhost:8080/test -H "Authorization: bearer <retrieved token from azure: eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1....."
```



## Rest Calls example ( cURL )
### Upload
curl https://graph.microsoft.com/v1.0/me/drive/root:/document1.docx:/content -X PUT -d @document1.docx -H "Authorization: bearer access_token_here"
###  Download
curl https://graph.microsoft.com/v1.0/me/drive/root:/document1.docx:/content -X GET  -H "Authorization: bearer access_token_here"
### Convert to PDF
curl https://graph.microsoft.com/v1.0/me/drive/root:/document1.docx:/content?format=pdf -o document1.pdf -H "Authorization: bearer access_token_here"
### Drive metadata
curl https://graph.microsoft.com/v1.0/drives/me -X GET  -H "Authorization: bearer access_token_here"