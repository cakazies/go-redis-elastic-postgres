FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
# Create appuser.
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/github.com/cakazies/go-redis-elastic-postgres
COPY . .
# Using go get.
RUN go get

# buildin apps in go-redis-elastic-postgres
RUN go build -o go-redis-elastic-postgres

# running go-redis-elastic-postgres
ENTRYPOINT ./go-redis-elastic-postgres

# running in port 
EXPOSE 7004
