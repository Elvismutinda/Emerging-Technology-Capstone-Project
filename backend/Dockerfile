FROM golang:1.23.2 AS builder

WORKDIR /app

COPY . .

RUN go mod download

# install the dependencies
RUN go get -v ./...

RUN go mod tidy

RUN go build -o app/backend/commands/server ./commands/server/main.go

EXPOSE 8000

# Start the application by running the compiled binary
ENTRYPOINT ["app/backend/commands/server"]