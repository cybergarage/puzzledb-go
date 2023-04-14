# Copyright (C) 2022 The PuzzleDB Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := bash

PKG_NAME=puzzledb
PKG_VER=$(shell git describe --abbrev=0 --tags)

MODULE_ROOT=github.com/cybergarage/puzzledb-go

PKG_SRC_ROOT=${PKG_NAME}
PKG=\
	${MODULE_ROOT}/${PKG_SRC_ROOT}/...

TEST_SRC_ROOT=${PKG_NAME}test
TEST_PKG=\
	${MODULE_ROOT}/${TEST_SRC_ROOT}/...

BIN_SRC_ROOT=bin
BIN_ID=${MODULE_ROOT}/${BIN_SRC_ROOT}
BIN_SERVER=${PKG_NAME}-server
BIN_DOCKER_TAG=cybergarage/${PKG_NAME}:${PKG_VER}
BIN_SERVER_ID=${BIN_ID}/${BIN_SERVER}
BIN_SRCS=\
        ${BIN_SRC_ROOT}/${BIN_SERVER}
BINS=\
        ${BIN_SERVER_ID}

.PHONY: test format vet lint clean docker

all: test

version:
	@pushd ${PKG_SRC_ROOT} && ./version.gen > version.go && popd

format:
	gofmt -s -w ${PKG_SRC_ROOT} ${TEST_SRC_ROOT} ${BIN_SRC_ROOT}

vet: format
	go vet ${PKG} ${TEST_PKG}

lint: format
	golangci-lint run ${PKG_SRC_ROOT}/... ${TEST_SRC_ROOT}/...

test: lint
	go test -v -p 1 -cover -timeout 60s ${PKG}/... ${TEST_PKG}/...

test_only:
	go test -v -p 1 -cover -timeout 60s ${PKG}/... ${TEST_PKG}/...

image:
	docker image build -t${BIN_DOCKER_TAG} .

build:
	go build -v -gcflags=${GCFLAGS} ${BINS}

install:
	go install -v -gcflags=${GCFLAGS} ${BINS}

run: build
	./${BIN_SERVER}

rund:
	docker container run -it --rm -p 6379:6379 -p 27017:27017 -p 3306:3306 ${BIN_DOCKER_TAG}

clean:
	go clean -i ${PKG}

#
# Document
#

%.md : %.adoc
	asciidoctor -b docbook -a leveloffset=+1 -o - $< | pandoc -t markdown_strict --wrap=none -f docbook > $@
csvs := $(wildcard doc/**/*.csv)
docs := $(patsubst %.adoc,%.md,$(wildcard *.adoc doc/*.adoc))
doc_touch: $(csvs)
	touch doc/*.adoc
doc: doc_touch $(docs)
	@mv README_.md README.md
	@sed -i '' -e "s/(img\//(doc\/img\//g" README.md

#
# Testing
#

%.pict : %.mod
	pict $< > $@
models=$(shell find ${TEST_SRC_ROOT} -name '*.mod')
picts=$(models:.mod=.pict)
pict: $(picts)
