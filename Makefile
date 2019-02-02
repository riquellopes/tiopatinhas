build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/tiopatinhas main.go
