# Start from golang base image
FROM golang:1.18-alpine3.15 AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app
#
RUN go install github.com/githubnemo/CompileDaemon@v1.4.0
#
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

## Command to run the executable
CMD /wait \
  && go run db/migrations/entry.go \
  && CompileDaemon --build="go build -o /app/cmd/api/api /app/cmd/api/api.go"  --command="/app/cmd/api/api" --color
