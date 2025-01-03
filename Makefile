SRC := cmd/server/main.go
TEMPL := github.com/a-h/templ/cmd/templ@latest
win32:
	@GOOS=windows go build -o bin/server_$@.exe $(SRC)
linux:
	@GOOS=linux go build -o bin/server_$@ $(SRC)
darwin:
	@GOOS=darwin go build -o bin/server_$@ $(SRC)
watch:
	@go run $(TEMPL) fmt . && go run $(TEMPL) generate --watch
act:
	@act -s GITHUB_TOKEN="$(shell gh auth token)"
update:
	@go get -u ./... && go mod tidy
format:
	@go fmt ./...
run:
	@go run $(SRC) $(filter-out $@,$(MAKECMDGOALS))
all: win32 linux darwin
%:
	@: