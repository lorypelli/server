SRC := cmd/server/main.go
win32:
	@GOOS=windows go build -o bin/server_$@.exe $(SRC)
linux:
	@GOOS=linux go build -o bin/server_$@ $(SRC)
darwin:
	@GOOS=darwin go build -o bin/server_$@ $(SRC)
watch:
	@go run github.com/a-h/templ/cmd/templ@latest fmt . && go run github.com/a-h/templ/cmd/templ@latest generate -watch
act:
	@act -s GITHUB_TOKEN="$(shell gh auth token)"
update:
	@go get -u ./... && go mod tidy
run:
	@go fmt ./... && go run $(SRC)
all: win32 linux darwin