Application = cli

.PHONY: build release clean run test


run:
	go run main.go

test:
	go test -v ./...

build:
	go build -o $(Application)

release:
	go build -ldflags "-s -w" -o $(Application)



clean:
	rm -f $(Application)
	rm -f $(Application).exe

