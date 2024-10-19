windows:
	@GOOS=windows go build -o bin/server_win32.exe main.go start.go changes.go
linux:
	@GOOS=linux go build -o bin/server_linux main.go start.go changes.go
darwin:
	@GOOS=darwin go build -o bin/server_darwin main.go start.go changes.go
build:
	@go build -o server.exe main.go start.go changes-go
run:
	@go run main.go start.go changes.go
all: windows linux darwin