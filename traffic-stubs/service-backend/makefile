MAKE_DIR = $(PWD)

test:
	go test

clean:

build:
	mkdir -p ${BUILD_DIR}
	go build -o ${BUILD_DIR}/service-backend service.go config.go
