
watch:
	goconvey -cover -excludedDirs testdata

test:
	go test -v

bench:
	go test -bench=. -run=XXX

doc:
	@echo "Doc server address: http://localhost:6060"
	godoc -http=":6060" -goroot=$$GOPATH
