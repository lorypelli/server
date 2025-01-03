FROM golang:alpine
WORKDIR /app
COPY . /app
EXPOSE 53273
RUN apk add --no-cache make
RUN make linux
ENTRYPOINT ["./bin/server_linux"]