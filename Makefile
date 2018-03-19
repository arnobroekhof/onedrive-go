dependencies:
	dep ensure -v

test:
	go test -v -cover -race ./onedrive ./auth
