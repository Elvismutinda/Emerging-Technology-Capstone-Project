
# Change this if necessary. I just wrote it to fill it up for now
# Delete this comment when you're done

FROM golang:1.23.2 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o backend .

EXPOSE 8000

ENTRYPOINT ["go", "run", "commands/server/main.go"]