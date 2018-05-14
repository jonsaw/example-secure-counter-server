build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/counterRead functions/counterRead/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/counterIncrement functions/counterIncrement/main.go
