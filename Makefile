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

#PREFIX?=$(shell pwd)
#GOPATH:=$(shell pwd)
#export GOPATH

PACKAGE_NAME=mimicdb

MODULE_ROOT=github.com/cybergarage/mimicdb
PACKAGE_ROOT=${MODULE_ROOT}/${PACKAGE_NAME}

SOURCE_ROOT=${PACKAGE_NAME}
TEST_SOURCE_ROOT=${PACKAGE_NAME}test
SOURCES=\
	${SOURCE_ROOT} \
	${TEST_SOURCE_ROOT}/plugins

PACKAGE_ID=${PACKAGE_ROOT}
PACKAGES=\
	${PACKAGE_ID} \
	${PACKAGE_ID}/obj \
	${PACKAGE_ID}/query \
	${PACKAGE_ID}/store \
	${PACKAGE_ID}/plugins \
	${PACKAGE_ID}/plugins/executor/llvm \
	${PACKAGE_ID}/plugins/executor/vdbe \
	${PACKAGE_ID}/plugins/query/mysql \
	${PACKAGE_ID}/plugins/store/memdb

.PHONY: version clean

all: test

format:
	gofmt -w ${SOURCES}

vet: format
	go vet ${PACKAGE_ROOT}

lint: format
	golangci-lint run ${SOURCES}

test: 
	go test -v -cover -timeout 60s ${PACKAGES}

clean:
	go clean -i ${PACKAGES}
