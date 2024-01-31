build:
	env GOOS=linux GOARCH=amd64 go build -v -o quicklinks_0.1-1/usr/local/bin .
	
lint:
	golangci-lint run -c ./golangci.yml ./...

test:
	go test ./... -v --cover

test-report:
	go test ./... -v --cover -coverprofile=coverage.out
	go tool cover -html=coverage.out

