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

PACKAGE_NAME=puzzledb

MODULE_ROOT=github.com/cybergarage
PACKAGE_ROOT=${MODULE_ROOT}/${PACKAGE_NAME}

SOURCE_ROOTS=\
	document \
	query \
	store \
	server \
	test

PACKAGE_ID=${PACKAGE_ROOT}

PACKAGES=\
	${PACKAGE_ID}/errors/... \
	${PACKAGE_ID}/document/... \
	${PACKAGE_ID}/query/... \
	${PACKAGE_ID}/store/... \
	${PACKAGE_ID}/server/...

TEST_PACKAGES=\
	${PACKAGE_ID}/test/...

.PHONY: test format vet lint clean

all: test

format:
	gofmt -w ${SOURCE_ROOTS}

vet: format
	go vet ${PACKAGE_ROOT}

lint: format
	golangci-lint run ${SOURCE_ROOTS}

test: 
	go test -v -cover -timeout 60s ${PACKAGES} ${TEST_PACKAGES}

clean:
	go clean -i ${PACKAGES}
