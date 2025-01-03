FROM golang:alpine
COPY . .
EXPOSE 53273
RUN apk add --no-cache make
RUN make linux
ENTRYPOINT ["./bin/server_linux"]