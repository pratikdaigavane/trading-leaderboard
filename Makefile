all:
	CGO_ENABLED=0 go build -o leaderboard -ldflags "-s" cmd/*.go

run:
	CGO_ENABLED=0 go run -ldflags "-s" cmd/*.go

test:
	go test ./...