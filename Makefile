check: test lint

test:
	go test ./...

lint:
	golangci-lint run

bench:
	go test ./... -bench=.