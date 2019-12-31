GOFMT=gofmt
GC=go build
VERSION := v0.0.1b-$(shell git describe --always --tags --long)
BUILD_NODE_PAR = -ldflags "-X github.com/TesraSupernet/Tesra/common/config.Version=$(VERSION)" #-race

SRC_FILES = $(shell git ls-files | grep -e .go$ | grep -v _test.go)

tesranode: $(SRC_FILES)
	$(GC)  $(BUILD_NODE_PAR) -o tesranode main.go
 
all: tesranode

tesranode-cross: tesranode-windows tesranode-linux tesranode-darwin

tesranode-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o tesranode-windows-amd64.exe main.go

tesranode-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o tesranode-linux-amd64 main.go

tesranode-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o tesranode-darwin-amd64 main.go

all-cross: tesranode-cross

format:
	$(GOFMT) -w main.go

clean:
	rm -rf *.8 *.o *.out *.6 *exe coverage
	rm -rf tesranode tesranode-*
