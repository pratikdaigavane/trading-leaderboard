.PHONY: build
build:
	CGO_ENABLED=0 go build -o leaderboard -ldflags "-s" cmd/*.go

.PHONY: run
run:
	CGO_ENABLED=0 go run -ldflags "-s" cmd/*.go

.PHONY: test
test:
	go test -v ./...