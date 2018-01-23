all: build

build:
	cd bin/; go build -a -installsuffix cgo -o ../fake_component_main_example .; cd ..

.PHONY: build
