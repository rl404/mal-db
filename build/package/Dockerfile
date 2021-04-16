# Golang base image
FROM golang:1.15 as go_builder
LABEL stage="mal_db_builder"
WORKDIR /go/src/github.com/rl404/mal-db
COPY . .
WORKDIR /go/src/github.com/rl404/mal-db/cmd/mal
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -mod vendor -o mal

# New stage from scratch
FROM alpine:3.13
RUN apk add --no-cache ca-certificates
COPY --from=go_builder /go/src/github.com/rl404/mal-db/cmd/mal/mal /cmd/mal/mal
COPY --from=go_builder /go/src/github.com/rl404/mal-db/fonts/Roboto-Bold.ttf /fonts/Roboto-Bold.ttf
WORKDIR /cmd/mal