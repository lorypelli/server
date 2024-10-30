SRC := ${wildcard *.go}
win32:
	@GOOS=windows go build -o bin/server_$@.exe ${SRC}
linux:
	@GOOS=linux go build -o bin/server_$@ ${SRC}
darwin:
	@GOOS=darwin go build -o bin/server_$@ ${SRC}
run:
	@go run ${SRC}
all: win32 linux darwin