Application = cli

.PHONY: build release clean run test pprof


run:
	go run main.go

test:
	go test -v ./...

build:
	go build -o $(Application)

release:
	go build -ldflags "-s -w" -o $(Application)

cross:
	set GOOS=linux
	set GOARCH=amd64
	go build -ldflags "-s -w"
	set GOOS=windows
	go build -ldflags "-s -w"

#学会获取参数才是真的
cpu:
	 go tool pprof -http=:8080 cpu.pprof
proxy:
	go env -w  GOPROXY=https://goproxy.io,direct


pprof:
	go tool pprof -http=:8080 *.pprof
clean:
	rm -f $(Application)
	rm -f $(Application).exe
	rm -f cpu.pprof

