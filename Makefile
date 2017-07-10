BINARY=chkit
PACKAGE = github.com/kfeofantov/chkit-v2
COMMIT_HASH = `git rev-parse --short HEAD 2>/dev/null`
BUILD_DATE = `date +%FT%T%Z`
LDFLAGS = -ldflags "-X ${PACKAGE}/chlib.CommitHash=${COMMIT_HASH} -X ${PACKAGE}/chlib.BuildDate=${BUILD_DATE}"


all: build

build:
	go build ${LDFLAGS} -o ${BINARY}

clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi
test:
	go test

build-all: clean-build
	mkdir build
	#Build for linux_386
	mkdir build/linux_386
	GOOS=linux GOARCH=386 go build -o build/linux_386/chkit
	#Build for linux_amd64
	mkdir build/linux_amd64
	GOOS=linux GOARCH=amd64 go build -o build/linux_amd64/chkit
	#Build for darwin_amd64
	mkdir build/darwin_amd64
	GOOS=darwin GOARCH=amd64 go build -o build/darwin_amd64/chkit
	#Build for windows_386
	#Build for windows_amd64

clean-build:
	rm -rf build

.PHONY: all build clean
