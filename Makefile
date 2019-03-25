VERSION:=$(shell git describe --tags --always --dirty --match='v*' 2> /dev/null || echo v0)
GO:=$(shell which go)

.PHONY: all build clean hive hive-server hive-game version

all: build

build: hive hive-server

clean:
	rm -rf build/

hive: build/hive/ hive-server hive-game
	cd hive/ && cp start.sh ../build/hive/ \

build/hive/:
	mkdir -p build/hive/

hive-server:
	cd cmd/server/ && GO111MODULE=on $(GO) build -ldflags="-s -w" -i -o server ./... \
	&& mv server ../../build/hive/ \
	&& cp -r www ../../build/hive/ \
	&& cp start.sh ../../build/hive/

hive-game:
	cd cmd/game/ && GOOS=js GOARCH=wasm GO111MODULE=on $(GO) build -ldflags="-s -w" -i -o ../../build/hive/www/game.wasm

version:
	@echo version: $(VERSION)

