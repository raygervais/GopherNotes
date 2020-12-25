build:
	go build -o bin/gn cmd/gophernotes/main.go

run:
	go run cmd/gophernotes/main.go

fmt:
	go fmt ./...

publish:
	GOOS=freebsd GOARCH=386 go build -o bin/gophernotes-freebsd-386 cmd/gophernotes/main.go
	GOOS=linux   GOARCH=386 go build -o bin/gophernotes-linux-386   cmd/gophernotes/main.go
	GOOS=windows GOARCH=386 go build -o bin/gophernotes-windows-386 cmd/gophernotes/main.go
