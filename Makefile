
watch:
	goconvey -cover -excludedDirs testdata

test:
	go test -v

bench:
	go test -bench=. -run=XXX

doc:
	godoc -http=":6060" -goroot=$$GOPATH
