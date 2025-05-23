container_runtime := $(shell which docker || which podman)
$(info using ${container_runtime})
up: down
	${container_runtime} compose up --build -d
down:
	${container_runtime} compose down
clean:
	${container_runtime} compose down -v
# will be...
run-tests:
	${container_runtime} run --rm --network=host tests:latest
test:
	make clean
	make up
	@echo wait cluster to start && sleep 2
	make run-tests
	make clean
	@echo "test finished"
tools:
	make -C bank-services tools
lint: tools
	make -C bank-services lint
proto:
	make -C bank-services protobuf