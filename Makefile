.PHONY: build

local-build: clean
	CC=/usr/bin/musl-gcc GOPRIVATE="github.com/opengovern" GOOS=linux GOARCH=amd64 go build -a -v -mod=mod -ldflags "-linkmode external -extldflags '-static' -s -w" -tags musl -o ./local/og-describer-googleworkspace main.go

build-cli: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -ldflags "-w -extldflags -static" -o ./build/og-googleworkspace-cli ./command/main.go

clean:
	rm -rf ./local/og-describer-googleworkspace ./build/og-googleworkspace-cli
