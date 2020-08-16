MAKE_DIR = $(PWD)
BUILD_DIR = $(PWD)/bin
.DEFAULT_GOAL := build

export BUILD_DIR

test:
	@$(MAKE) -C testr-envs test
	@$(MAKE) -C testr-runners test
	@$(MAKE) -C traffic-proxy test
	@$(MAKE) -C traffic-stubs test
	@$(MAKE) -C traffic-storage test

clean:
	rm -rf $(BUILD_DIR)
	@$(MAKE) -C testr-envs clean
	@$(MAKE) -C testr-runners clean
	@$(MAKE) -C traffic-proxy clean
	@$(MAKE) -C traffic-stubs clean
	@$(MAKE) -C traffic-storage clean

build:
	@$(MAKE) -C testr-envs build
	@$(MAKE) -C testr-runners build
	@$(MAKE) -C traffic-proxy build
	@$(MAKE) -C traffic-stubs build
	@$(MAKE) -C traffic-storage build

package:

compile:
	@$(MAKE) build
