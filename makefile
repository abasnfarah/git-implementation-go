build:
	go build -o bin/mygit cmd/mygit/main.go

run:
	go run cmd/mygit/main.go

clean:
	rm -rf bin

test:
	go test -v ./...
