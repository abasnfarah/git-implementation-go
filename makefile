build:
	go build -o ~/.local/bin/mygit cmd/mygit/main.go

run:
	go run cmd/mygit/main.go

clean:
	rm -rf ~/.local/bin/mygit

test:
	go test -v ./...
