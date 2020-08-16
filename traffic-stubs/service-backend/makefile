MAKE_DIR = $(PWD)

test:
	go test

clean:

build:
	mkdir -p ${BUILD_DIR}
	cp config.yaml ${BUILD_DIR}/service-backend.yaml
	go build -o ${BUILD_DIR}/service-backend service.go

package:
