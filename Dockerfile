FROM golang:alpine AS build
ARG TARGETARCH
WORKDIR /src
COPY . .
RUN apk add --no-cache make && make generate linux ARCH=$TARGETARCH
FROM scratch
COPY --from=build /src/bin/server_linux /server
WORKDIR /app
EXPOSE 53273
ENTRYPOINT ["/server"]