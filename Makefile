SHELL = /bin/bash

all:
	@echo goboom makefile
	@echo \"make build\" to build

.PHONY: build
build:
	go build goboom.go

.PHONY: package
package: build
	tar cf goboom_linux_amd64.tar.xz goboom goboom_run README.rst docs/_build/man/goboom.1
	sha512sum goboom_linux_amd64.tar.xz > goboom_linux_amd64.tar.xz.sha512sum

.PHONY: html
html:
	pushd docs && make html && popd

.PHONY: man
man:
	pushd docs && make man && popd
