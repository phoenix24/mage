MAKE_DIR = $(PWD)

test:
	go test

clean:

build:
	mkdir -p $(BUILD_DIR)
	cp config.yaml $(BUILD_DIR)/traffic-proxy.yaml
	go build -o $(BUILD_DIR)/traffic-proxy server.go

package:
