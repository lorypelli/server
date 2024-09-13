windows:
	@GOOS=windows go build -o bin/server_win32.exe main.go start.go
linux:
	@GOOS=linux go build -o bin/server_linux main.go start.go
darwin:
	@GOOS=darwin go build -o bin/server_darwin main.go start.go
build:
	@go build -o server.exe main.go start.go
run:
	@go run main.go start.go
all: windows linux darwin