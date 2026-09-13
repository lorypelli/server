SRC := cmd/server/main.go
ARCH ?= amd64
BUILD := CGO_ENABLED=0 go build -trimpath -ldflags="-s -w"
all: win32 linux darwin linux_arm64 darwin_arm64
win32:
	@GOOS=windows GOARCH=$(ARCH) $(BUILD) -o bin/server_$@.exe $(SRC)
linux:
	@GOOS=linux GOARCH=$(ARCH) $(BUILD) -o bin/server_$@ $(SRC)
darwin:
	@GOOS=darwin GOARCH=$(ARCH) $(BUILD) -o bin/server_$@ $(SRC)
linux_arm64:
	@GOOS=linux GOARCH=arm64 $(BUILD) -o bin/server_$@ $(SRC)
darwin_arm64:
	@GOOS=darwin GOARCH=arm64 $(BUILD) -o bin/server_$@ $(SRC)
generate:
	@go tool templ generate
act:
	@act -s GITHUB_TOKEN="$(shell gh auth token)"
update:
	@go get -u ./... && go mod tidy
run:
	@go run $(SRC) $(filter-out $@,$(MAKECMDGOALS))
docker:
	@docker build . -t server
%:
	@:
