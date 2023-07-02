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


windows:
	set GOOS=windows
	go build -ldflags "-s -w"
linux:
	set GOOS=linux
	go build -ldflags "-s -w"
cross: linux windows
	@echo "cross build done"

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

