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

MODULE_ROOT=github.com/cybergarage/puzzledb-go

PKG_SRC_ROOT=${PKG_NAME}
PKG=\
	${MODULE_ROOT}/${PKG_SRC_ROOT}/...

TEST_SRC_ROOT=${PKG_NAME}test
TEST_PKG=\
	${MODULE_ROOT}/${TEST_SRC_ROOT}/...

TEST_PKG=${MODULE_ROOT}/${TEST_SRC_ROOT}

.PHONY: test format vet lint clean

all: test

format:
	gofmt -s -w ${PKG_SRC_ROOT} ${TEST_SRC_ROOT}

vet: format
	go vet ${PKG}

lint: format
	golangci-lint run ${PKG_SRC_ROOT}/... ${TEST_SRC_ROOT}/...

test:
	go test -v -cover -timeout 60s ${PKG}/... ${TEST_PKG}/...

clean:
	go clean -i ${PKG}
