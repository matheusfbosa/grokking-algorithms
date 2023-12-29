run:
	go run example.go

test:
	go test ./...

bench:
	go test -bench=. -benchmem ./...
