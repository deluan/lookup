
watch:
	goconvey -cover -excludedDirs examples

test:
	go test ./...

bench:
	go test -bench=. -run=XXX

doc:
	godoc -http=":6060" -goroot=$$GOPATH
