MAKE_DIR = $(PWD)

test:
	go test

clean:

build:
	mkdir -p ${BUILD_DIR}
	cp config.yaml ${BUILD_DIR}/service-frontend.yaml
	go build -o ${BUILD_DIR}/service-frontend service.go

package:
