SRC := cmd/server/main.go
win32:
	@GOOS=windows go build -o bin/server_$@.exe ${SRC}
linux:
	@GOOS=linux go build -o bin/server_$@ ${SRC}
darwin:
	@GOOS=darwin go build -o bin/server_$@ ${SRC}
watch:
	@templ generate -watch
run:
	@go fmt all && go run ${SRC}
start:
	@make -j 2 watch run