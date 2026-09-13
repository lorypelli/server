SRC := ./cmd/server
BUILD := CGO_ENABLED=0 go build -trimpath -ldflags="-s -w"
all: linux_amd64 linux_arm64 windows_amd64 windows_arm64 darwin_amd64 darwin_arm64
linux_amd64:
	@GOOS=linux GOARCH=amd64 $(BUILD) -o bin/server_linux_amd64 $(SRC)
linux_arm64:
	@GOOS=linux GOARCH=arm64 $(BUILD) -o bin/server_linux_arm64 $(SRC)
windows_amd64:
	@GOOS=windows GOARCH=amd64 $(BUILD) -o bin/server_windows_amd64.exe $(SRC)
windows_arm64:
	@GOOS=windows GOARCH=arm64 $(BUILD) -o bin/server_windows_arm64.exe $(SRC)
darwin_amd64:
	@GOOS=darwin GOARCH=amd64 $(BUILD) -o bin/server_darwin_amd64 $(SRC)
darwin_arm64:
	@GOOS=darwin GOARCH=arm64 $(BUILD) -o bin/server_darwin_arm64 $(SRC)
generate:
	@go tool templ generate
act:
	@act -s GITHUB_TOKEN="$(shell gh auth token)"
update:
	@go get -u ./... && go mod tidy
run:
	@go run $(SRC) $(filter-out $@,$(MAKECMDGOALS))
docker:
	@docker build . -t server && docker run -it --rm -p 53273:53273 server
%:
	@:
