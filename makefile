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

build-testr-envs:
	@$(MAKE) -C testr-envs build

build-testr-runners:
	@$(MAKE) -C testr-runners build

build-traffic-proxy:
	@$(MAKE) -C traffic-proxy build

build-traffic-stubs:
	@$(MAKE) -C traffic-stubs build

build-traffic-storage:
	@$(MAKE) -C traffic-storage build

build:
	@echo "build everything!"
	@$(MAKE) build-testr-envs
	@$(MAKE) build-testr-runners
	@$(MAKE) build-traffic-proxy
	@$(MAKE) build-traffic-stubs
	@$(MAKE) build-traffic-storage

package:

compile:
	@$(MAKE) build
