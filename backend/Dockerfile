
# Change this if necessary. I just wrote it to fill it up for now
# Delete this comment when you're done

FROM golang:1.23.2 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/backend commands/server/main.go

# Expose the port the app will run on
EXPOSE 8000

# Start the application by running the compiled binary
ENTRYPOINT ["/app/backend"]