SRC := cmd/server/main.go
win32:
	@GOOS=windows go build -o bin/server_$@.exe ${SRC}
linux:
	@GOOS=linux go build -o bin/server_$@ ${SRC}
darwin:
	@GOOS=darwin go build -o bin/server_$@ ${SRC}
run:
	@templ generate && go run ${SRC}
all: win32 linux darwin