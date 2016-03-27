SHELL = /bin/bash
VERSION=0.2.1

all:
	@echo goboom makefile $(VERSION)
	@echo \"make build\" to build

.PHONY: build
build:
	goxc -pv="$(VERSION)" -build-ldflags='-X main.version=$(VERSION)' xc

.PHONY: package
package: man
	goxc -pv="$(VERSION)" -build-ldflags='-X main.version=$(VERSION)'

.PHONY: html
html:
	pushd docs && make html SPHINXOPTS="-Dversion=$(VERSION) -Drelease=$(VERSION)" && popd

.PHONY: man
man:
	pushd docs && make man SPHINXOPTS="-Dversion=$(VERSION) -Drelease=$(VERSION)" && popd

.PHONY: upload
upload:
	goxc -pv="$(VERSION)" bintray
