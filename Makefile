
watch:
	goconvey -cover -excludedDirs examples

test:
	go test ./...

doc:
	godoc -http=":6060"
