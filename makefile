VER = 0.2.0
SRC = cmd/hashi.up/hashi.up.go

all: test build

.PHONY: test
test:
	@go test -v ./...

.PHONY: build
build:
	@go build -o dist/hashi.up ${SRC}

.PHONY: install
install:
	@mv dist/hashi.up /usr/local/bin && echo "hashi.up has been installed to /usr/local/bin/hashi.up. Please update your path to reflect this location if is isn't there already."

.PHONY: build_all
build_all: build_nix build_windows

.PHONY: build_nix
build_nix: nix_amd64 nix_386 nix_arm64 nix_arm

.PHONY: nix_amd64
nix_amd64:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o dist/nix/hashi.up ${SRC}
	@cd dist/nix; tar -czf hashi.up_${VER}_linux_amd64.tar.gz ./hashi.up; cd ../..
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o dist/nix/hashi.up ${SRC}
	@cd dist/nix; tar -czf hashi.up_${VER}_darwin_amd64.tar.gz ./hashi.up; cd ../..
	@rm -f dist/nix/hashi.up

.PHONY: nix_386
nix_386:
	@CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -o dist/nix/hashi.up ${SRC}
	@cd dist/nix; tar -czf hashi.up_${VER}_linux_386.tar.gz ./hashi.up; cd ../..
	@rm -f dist/nix/hashi.up

.PHONY: nix_arm64
nix_arm64:
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -o dist/nix/hashi.up ${SRC}
	@cd dist/nix; tar -czf hashi.up_${VER}_linux_arm64.tar.gz ./hashi.up; cd ../..
	@rm -f dist/nix/hashi.up

.PHONY: nix_arm
nix_arm:
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -o dist/nix/hashi.up ${SRC}
	@cd dist/nix; tar -czf hashi.up_${VER}_linux_arm.tar.gz ./hashi.up; cd ../..
	@rm -f dist/nix/hashi.up

.PHONY: build_windows
build_windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -o dist/win/hashi.up.exe ${SRC}
	@cd dist/win; zip hashi.up_${VER}_windows_amd64.zip ./hashi.up.exe; cd ../..
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -a -o dist/win/hashi.up.exe ${SRC}
	@cd dist/win; zip hashi.up_${VER}_windows_386.zip ./hashi.up.exe; cd ../..
	@rm -f dist/win/hashi.up.exe

