migrate:
	go run ./cmd/cli/main.go
run:
	go run ./cmd/server/main.go

build:
	go build ./cmd/server/main.go

watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -s -r '\.go$$' go run ./cmd/server/main.go
