FROM golang:alpine AS build
ARG TARGETARCH
WORKDIR /src
COPY . .
RUN apk add --no-cache make
RUN make linux_$TARGETARCH
FROM scratch
ARG TARGETARCH
COPY --from=build /src/bin/server_linux_$TARGETARCH /server
WORKDIR /app
EXPOSE 53273
ENTRYPOINT ["/server"]