.PHONY: test fmt

test:
	go test -v -cover ./...

gen:
	GENERATE_SVG=true go test ./...

fmt:
	go mod tidy
	gofmt -s -w .
	goarrange run -r .
	golangci-lint run ./...
