FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./bin/contact_book cmd/api/main.go

CMD ["./bin/contact_book"]
