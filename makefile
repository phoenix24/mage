MAKE_DIR = $(PWD)
BUILD_DIR = $(PWD)/bin
.DEFAULT_GOAL := build

export BUILD_DIR

test:

clean:
	rm -rf $(BUILD_DIR)
	@$(MAKE) -C testr-envs clean
	@$(MAKE) -C testr-runners clean
	@$(MAKE) -C traffic-proxy clean
	@$(MAKE) -C traffic-stubs clean
	@$(MAKE) -C traffic-sniffer clean
	@$(MAKE) -C traffic-storage clean

build:
	@$(MAKE) -C testr-envs build
	@$(MAKE) -C testr-runners build
	@$(MAKE) -C traffic-proxy build
	@$(MAKE) -C traffic-stubs build
	@$(MAKE) -C traffic-sniffer build
	@$(MAKE) -C traffic-storage build

compile:
	@$(MAKE) build
