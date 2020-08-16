MAKE_DIR = $(PWD)

test:
	@$(MAKE) -C service-backend test
	@$(MAKE) -C service-frontend test

clean:

build:
	mkdir -p $(BUILD_DIR)
	@$(MAKE) -C service-backend build
	@$(MAKE) -C service-frontend build
