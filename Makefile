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

GOBIN := $(shell go env GOPATH)/bin
PATH := $(GOBIN):$(PATH)

PKG_NAME=puzzledb
PKG_VER=$(shell git describe --abbrev=0 --tags)
PKG_COVER=${PKG_NAME}-cover

MODULE_ROOT=github.com/cybergarage/puzzledb-go
PKG_ROOT=${MODULE_ROOT}/${PKG_NAME}

PKG_SRC_ROOT=${PKG_NAME}
PKG=\
	${MODULE_ROOT}/${PKG_SRC_ROOT}/...

TEST_SRC_ROOT=${PKG_NAME}test
TEST_PKG=\
	${MODULE_ROOT}/${TEST_SRC_ROOT}/...

BIN_SRC_ROOT=cmd
BIN_ID=${MODULE_ROOT}/${BIN_SRC_ROOT}
BIN_SERVER=${PKG_NAME}-server
BIN_SERVER_ID=${BIN_ID}/${BIN_SERVER}
BIN_SERVER_DOCKER_TAG=cybergarage/${PKG_NAME}:${PKG_VER}
BIN_SERVER_DOCKER_TAG_PRE=cybergarage/${PKG_NAME}:${PKG_VER}-pre
BIN_SERVER_DOCKER_TAG_LATEST=cybergarage/${PKG_NAME}:latest
BIN_CLI=${PKG_NAME}-cli
BIN_CLI_ID=${BIN_ID}/${BIN_CLI}
BIN_SRCS=\
        ${BIN_SRC_ROOT}/${BIN_SERVER} \
        ${BIN_SRC_ROOT}/${BIN_CLI}
BINS=\
        ${BIN_SERVER_ID} \
        ${BIN_CLI_ID}

.PHONY: test format vet lint clean docker cmd

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
	go test -v -p 1 -timeout 60s -cover -coverpkg=${PKG} -coverprofile=${PKG_COVER}.out ${PKG}/... ${TEST_PKG}/...
	go tool cover -html=${PKG_COVER}.out -o ${PKG_COVER}.html

test_only:
	go test -v -p 1 -timeout 60s -cover -coverpkg=${PKG} -coverprofile=${PKG_COVER}.out ${PKG}/... ${TEST_PKG}/...
	go tool cover -html=${PKG_COVER}.out -o ${PKG_COVER}.html

image:
	docker image build -t${BIN_SERVER_DOCKER_TAG} -t${BIN_SERVER_DOCKER_TAG_PRE} -t${BIN_SERVER_DOCKER_TAG_LATEST} .
	docker push ${BIN_SERVER_DOCKER_TAG_PRE}
	docker push ${BIN_SERVER_DOCKER_TAG_LATEST}

cmd:
	go build -v -gcflags=${GCFLAGS} ${BINS}

install:
	go install -v -gcflags=${GCFLAGS} ${BINS}

run: cmd
	./${BIN_SERVER}

rund:
	docker container run -it --rm -p 6379:6379 -p 27017:27017 -p 3306:3306 -p 50053:50053 -p 9181:9181 ${BIN_SERVER_DOCKER_TAG_LATEST}

log:
	git log ${PKG_VER}..HEAD --date=short --no-merges --pretty=format:"%s"

clean:
	go clean -i ${PKG}

watch:
	fswatch -o . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make test
#
# Document
#

DOC_CLI_ROOT=doc/cmd/cli
DOC_CLI_BIN=puzzledb-cli-doc
doc_cmd_cli:
	go build -o ${DOC_CLI_ROOT}/${DOC_CLI_BIN} ${MODULE_ROOT}/${DOC_CLI_ROOT}
	pushd ${DOC_CLI_ROOT} && ./${DOC_CLI_BIN} && popd
	git add ${DOC_CLI_ROOT}/*.md

DOC_SERVER_ROOT=doc/cmd/server
DOC_SERVER_BIN=puzzledb-server-doc
doc_cmd_server:
	go build -o ${DOC_SERVER_ROOT}/${DOC_SERVER_BIN} ${MODULE_ROOT}/${DOC_SERVER_ROOT}
	pushd ${DOC_SERVER_ROOT} && ./${DOC_SERVER_BIN} && popd
	git add ${DOC_SERVER_ROOT}/*.md

cmd_docs: doc_cmd_cli doc_cmd_server

%.md : %.adoc
	asciidoctor -b docbook -a leveloffset=+1 -o - $< | pandoc -t markdown_strict --wrap=none -f docbook > $@
csvs := $(wildcard doc/*/*.csv doc/*/*/*.csv)
docs := $(patsubst %.adoc,%.md,$(wildcard *.adoc doc/*.adoc doc/*/*.adoc))
doc_touch: $(csvs)
	touch doc/*.adoc doc/*/*.adoc

doc: doc_touch $(docs) cmd_docs
	@mv README_.md README.md
	@sed -i '' -e "s/(img\//(doc\/img\//g" README.md
	@rm ChangeLog.md

#
# Protos
#

PKG_PROTO_ROOT=${PKG_SRC_ROOT}/proto
protopkg:
	go get -u google.golang.org/protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest	
%.pb.go : %.proto
	protoc -I=${PKG_PROTO_ROOT} --go_out=paths=source_relative:${PKG_PROTO_ROOT}/grpc --go-grpc_out=paths=source_relative:${PKG_PROTO_ROOT}/grpc $<
protos=$(shell find ${PKG_SRC_ROOT} -name '*.proto')
pbs=$(protos:.proto=.pb.go)
proto: $(pbs)

#
# Testing
#

%.pict : %.mod
	pict $< > $@
models=$(shell find ${TEST_SRC_ROOT} -name '*.mod')
picts=$(models:.mod=.pict)
pict: $(picts)
