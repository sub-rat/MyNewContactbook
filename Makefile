run: build
	./bin/contact_book

build:
	go build -o ./bin/contact_book cmd/api/main.go

clean:
	go clean

test:
	go test ./...