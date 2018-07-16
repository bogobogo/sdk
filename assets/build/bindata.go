// Code generated by go-bindata.
// sources:
// etc/build/Makefile
// etc/build/etc/build/Dockerfile.build.alpine.tpl
// etc/build/etc/build/Dockerfile.build.debian.tpl
// etc/build/etc/run.sh
// etc/build/make/bootstrap.mk
// etc/build/make/functions.mk
// etc/build/make/help.mk
// etc/build/make/rules.mk
// DO NOT EDIT!

package build

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _makefile = []byte(`# This is the entrypoint of the bblfsh make SDK, the SDK should be stored at
# ` + "`" + `sdklocation` + "`" + `, generated executing ` + "`" + `bblfsh-sdk prepare-build` + "`" + `. Many different
# vars as ` + "`" + `LANGUAGE` + "`" + ` are required, this variables are extracting from the
# ` + "`" + `manifest.toml` + "`" + ` at the root and the project translating it to a make include
# at ` + "`" + `manifest` + "`" + ` using ` + "`" + `bblfsh-sdk manifest` + "`" + ` tool.

location = $(shell pwd)
sdklocation = $(location)/.sdk
manifest = $(sdklocation)/make/manifest.mk
tmplocation = $(sdklocation)/tmp
$(shell mkdir -p $(tmplocation))

bblfsh-sdk := $(shell command -v bblfsh-sdk 2> /dev/null)
bblfsh-sdk-tools := $(shell command -v bblfsh-sdk-tools 2> /dev/null)
in-container := $(shell echo $$ENVIRONMENT)
host-platform := $(shell echo $$HOST_PLATFORM)

ifdef bblfsh-sdk-tools # run only with Golang
    ifdef in-container
        bblfsh-sdk-tools :=  go run /go/src/gopkg.in/bblfsh/sdk.v2/cmd/bblfsh-sdk-tools/main.go
    endif

    $(shell $(bblfsh-sdk-tools) envvars > $(manifest))
endif

include $(sdklocation)/make/manifest.mk
include $(sdklocation)/make/functions.mk
include $(sdklocation)/make/bootstrap.mk
include $(sdklocation)/make/help.mk
include $(sdklocation)/make/rules.mk
sdkloaded = true
`)

func makefileBytes() ([]byte, error) {
	return _makefile, nil
}

func makefile() (*asset, error) {
	bytes, err := makefileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Makefile", size: 1196, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcBuildDockerfileBuildAlpineTpl = []byte(`FROM ${DOCKER_BUILD_NATIVE_IMAGE}

# remove any pre-installed Go SDK in the base image and reset GOROOT
RUN sh -c '[[ ! -z $(which go) ]] && rm -rf $(go env GOROOT) || true'
ENV GOROOT=""

ENV GOLANG_SRC_URL https://golang.org/dl/go${RUNTIME_GO_VERSION}.src.tar.gz

# from https://github.com/docker-library/golang/blob/master/1.9/alpine3.7/Dockerfile

RUN apk add --update --no-cache ca-certificates openssl && update-ca-certificates
RUN wget https://raw.githubusercontent.com/docker-library/golang/e63ba9c5efb040b35b71e16722b71b2931f29eb8/${RUNTIME_GO_VERSION}/alpine3.7/no-pic.patch -O /no-pic.patch -O /no-pic.patch

RUN set -ex \
	&& apk add --no-cache --virtual .build-deps \
		bash \
		gcc \
		musl-dev \
		openssl \
		go \
	\
	&& export GOROOT_BOOTSTRAP="$(go env GOROOT)" \
	\
	&& wget -q "$GOLANG_SRC_URL" -O golang.tar.gz \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz \
	&& cd /usr/local/go/src \
	&& patch -p2 -i /no-pic.patch \
	&& ./make.bash \
	\
	&& rm -rf /*.patch \
	&& apk del .build-deps

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
`)

func etcBuildDockerfileBuildAlpineTplBytes() ([]byte, error) {
	return _etcBuildDockerfileBuildAlpineTpl, nil
}

func etcBuildDockerfileBuildAlpineTpl() (*asset, error) {
	bytes, err := etcBuildDockerfileBuildAlpineTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/build/Dockerfile.build.alpine.tpl", size: 1156, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcBuildDockerfileBuildDebianTpl = []byte(`FROM ${DOCKER_BUILD_NATIVE_IMAGE}

# remove any pre-installed Go SDK in the base image and reset GOROOT
RUN sh -c '[[ ! -z $(which go) ]] && rm -rf $(go env GOROOT) || true'
ENV GOROOT=""

ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go${RUNTIME_GO_VERSION}.linux-amd64.tar.gz

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"`)

func etcBuildDockerfileBuildDebianTplBytes() ([]byte, error) {
	return _etcBuildDockerfileBuildDebianTpl, nil
}

func etcBuildDockerfileBuildDebianTpl() (*asset, error) {
	bytes, err := etcBuildDockerfileBuildDebianTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/build/Dockerfile.build.debian.tpl", size: 528, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcRunSh = []byte(`#!/bin/sh
darkcyan='\033[0;31m'
normal=$'\e[0m'

if [ $ENVIRONMENT ]; then
    ENVIRONMENT="[$ENVIRONMENT]"
fi

echo -e $darkcyan+ $ENVIRONMENT $@$normal

LOG=` + "`" + `$@ 2>&1` + "`" + `
RETVAL=$?
if [ $RETVAL -gt 0 ] || [ $VERBOSE ] ; then
    echo "$LOG"
fi

exit $RETVAL
`)

func etcRunShBytes() ([]byte, error) {
	return _etcRunSh, nil
}

func etcRunSh() (*asset, error) {
	bytes, err := etcRunShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/run.sh", size: 256, mode: os.FileMode(484), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _makeBootstrapMk = []byte(`# variables being included from the ` + "`" + `manifest.mk` + "`" + `
LANGUAGE ?=
RUNTIME_OS ?=
RUNTIME_NATIVE_VERSION ?=
RUNTIME_GO_VERSION ?=

# get the git commit
GIT_COMMIT=$(shell git rev-parse HEAD | cut -c1-7)
GIT_DIRTY=$(shell test -n "` + "`" + `git status --porcelain -uno` + "`" + `" && echo "-dirty" || true)

# optional variables
DRIVER_DEV_PREFIX := dev
DRIVER_VERSION ?= $(DRIVER_DEV_PREFIX)-$(GIT_COMMIT)$(GIT_DIRTY)

DOCKER_IMAGE ?= bblfsh/$(LANGUAGE)-driver
DOCKER_IMAGE_VERSIONED ?= $(call escape_docker_tag,$(DOCKER_IMAGE):$(DRIVER_VERSION))
DOCKER_BUILD_NATIVE_IMAGE ?= $(DOCKER_IMAGE)-build
DOCKER_BUILD_DRIVER_IMAGE ?= $(DOCKER_IMAGE)-build-with-go

# defined behaviour for builds inside travis-ci
ifneq ($(origin CI), undefined)
    # if we are inside CI, verbose is enabled by default
	VERBOSE := true
endif

# if TRAVIS_TAG defined DRIVER_VERSION is overrided
ifneq ($(TRAVIS_TAG), )
    DRIVER_VERSION := $(TRAVIS_TAG)
endif

# if we are not in master, and it's not a tag the push is disabled
ifneq ($(TRAVIS_BRANCH), master)
	ifeq ($(TRAVIS_TAG), )
        pushdisabled = "push disabled for non-master branches"
	endif
endif

# if this is a pull request, the push is disabled
ifneq ($(TRAVIS_PULL_REQUEST), false)
        pushdisabled = "push disabled for pull-requests"
endif
`)

func makeBootstrapMkBytes() ([]byte, error) {
	return _makeBootstrapMk, nil
}

func makeBootstrapMk() (*asset, error) {
	bytes, err := makeBootstrapMkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "make/bootstrap.mk", size: 1265, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _makeFunctionsMk = []byte(`# escape_docker_tag escape colon char to allow use a docker tag as rule
define escape_docker_tag
$(subst :,--,$(1))
endef

# unescape_docker_tag an escaped docker tag to be use in a docker command
define unescape_docker_tag
$(subst --,:,$(1))
endef`)

func makeFunctionsMkBytes() ([]byte, error) {
	return _makeFunctionsMk, nil
}

func makeFunctionsMk() (*asset, error) {
	bytes, err := makeFunctionsMkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "make/functions.mk", size: 248, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _makeHelpMk = []byte(`define helptext
bblfsh-sdk build system
=======================

The bblfsh build system helps to you to build and test a bblfsh driver. It
contains several rules to build docker containers, execute tests, validate
the driver, etc.

RULES
make build            builds driver's docker image, compiling the normalizer
                      component and the native component if needed calling the
                      rules: build-native and build-driver. Builds the required
                      docker images to do this.

make build-native     compiles the native component if needed, in interpreted
                      languages only prepares the scripts to execute the
                      component. To perform this is executes make calling the
                      private rule: ` + "`" + `build-native-internal` + "`" + ` defined in the
                      Makefile in the root of the project inside of the build
                      container.

make build-driver     compiles the normalizer component.

make test             execute all the unit tests of the components inside of the
                      build containers. It build the docker images if need it.

make test-native      execute the unit test for the native component. To perform
                      this is execute make calling the private rule:
                      ` + "`" + `test-native-internal` + "`" + ` defined in the Makefile in the root
                      of the project inside of the build container.

make test-driver      execute the unit test for the normalizer component.

make push             push the driver's docker image to the Docker registry. This
                      rule can be only executed inside of a Travis-CI environment
                      and just when is running for a tag.

make integration-test execute integration tests.

make validate         validates the current driver.

make clean            cleans all the build directories.

INTERNAL RULES
Two internal rules are required to run the test and the build main rules:
` + "`" + `test-native-internal` + "`" + ` and ` + "`" + `build-native-internal` + "`" + `. This rules are defined in the
Makefile in the root of the driver, they contain the language specific commands
for the native runtime.

To further documentation please go to https://doc.bblf.sh
endef

help:
	@echo "$$helptext" | more
`)

func makeHelpMkBytes() ([]byte, error) {
	return _makeHelpMk, nil
}

func makeHelpMk() (*asset, error) {
	bytes, err := makeHelpMkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "make/help.mk", size: 2291, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _makeRulesMk = []byte(`BUILD_PATH := $(location)/build

RUN := $(sdklocation)/etc/run.sh
RUN_VERBOSE := VERBOSE=1 $(RUN)

GRPC_PORT := 9432
GRPC_PORT_INTEGRATION ?= 39432

# docker runtime commands
DOCKER_CMD ?= docker
DOCKER_BUILD ?= $(DOCKER_CMD) build
DOCKER_RUN ?= $(DOCKER_CMD) run --rm -t
DOCKER_TAG ?= $(DOCKER_CMD) tag
DOCKER_PUSH ?= $(DOCKER_CMD) push
DOCKER_SOCK ?= /var/run/docker.sock

BUILD_VOLUME_TARGET ?= /opt/driver/src/
BUILD_VOLUME_PATH ?= $(location)

DOCKER_FILE_$(DOCKER_IMAGE_VERSIONED) ?= $(location)/Dockerfile.tpl
DOCKER_FILE_$(DOCKER_BUILD_DRIVER_IMAGE) ?= $(sdklocation)/etc/build/Dockerfile.build.$(RUNTIME_OS).tpl
DOCKER_FILE_$(DOCKER_BUILD_NATIVE_IMAGE) ?= $(location)/Dockerfile.build.tpl

# list of images to build
BUILD_IMAGE=$(DOCKER_BUILD_NATIVE_IMAGE) $(DOCKER_BUILD_DRIVER_IMAGE) $(DOCKER_IMAGE_VERSIONED)

# golang runtime commands
GO_CMD = go
GO_RUN = $(GO_CMD) run
GO_TEST = $(GO_CMD) test -v
# If GOPATH has multiple dir entries, pick the first one
GO_PATH := $(shell echo ${GOPATH}|cut -f1 -d':')

# build enviroment variables
BUILD_USER ?= bblfsh
BUILD_UID ?= $(shell id -u $(USER))
BUILD_ARGS ?=
BUILD_NATIVE_CMD ?= $(DOCKER_RUN) \
	-u $(BUILD_USER):$(BUILD_UID) \
	-v $(BUILD_VOLUME_PATH):$(BUILD_VOLUME_TARGET) \
	-v $(GO_PATH):/go \
	-e ENVIRONMENT=$(DOCKER_BUILD_NATIVE_IMAGE) \
	-e HOST_PLATFORM=$(shell uname) \
	$(DOCKER_BUILD_NATIVE_IMAGE)

BUILD_DRIVER_CMD ?= $(DOCKER_RUN) \
    -v $(DOCKER_SOCK):$(DOCKER_SOCK) \
	-u $(BUILD_USER):$(BUILD_UID) \
	-v $(BUILD_VOLUME_PATH):$(BUILD_VOLUME_TARGET) \
	-v $(GO_PATH):/go \
	-e ENVIRONMENT=$(DOCKER_BUILD_DRIVER_IMAGE) \
	-e HOST_PLATFORM=$(shell uname) \
	$(DOCKER_BUILD_DRIVER_IMAGE)

# if VERBOSE is unset docker build is executed in quite mode
ifeq ($(origin VERBOSE), undefined)
	BUILD_ARGS += -q
endif

ALLOWED_IN_DOCKERFILE = \
	LANGUAGE \
	RUNTIME_NATIVE_VERSION RUNTIME_GO_VERSION \
	BUILD_USER BUILD_UID BUILD_PATH \
	DOCKER_IMAGE DOCKER_IMAGE_VERSIONED DOCKER_BUILD_NATIVE_IMAGE

# we export the variable to allow envsubst, substitute the vars in the
# Dockerfiles
export

all: build

check-gopath:
ifeq ($(GOPATH),)
	$(error GOPATH is not defined)
endif

$(BUILD_PATH):
	@$(RUN) mkdir -p $(BUILD_PATH)/bin
	@$(RUN) mkdir -p $(BUILD_PATH)/etc

$(BUILD_IMAGE):
	@$(eval TEMP_FILE := "$(tmplocation)/tmp.$(shell date "+%s-%N")")
	@eval "envsubst '$(foreach v,$(ALLOWED_IN_DOCKERFILE),\$${$(v)})' < $(DOCKER_FILE_$@) > $(TEMP_FILE)"
	@$(RUN) $(DOCKER_BUILD) $(BUILD_ARGS) -t $(call unescape_docker_tag,$@) -f $(TEMP_FILE) .
	@rm $(TEMP_FILE)

test: | validate test-native test-driver
test-native: $(DOCKER_BUILD_NATIVE_IMAGE)
	@$(RUN_VERBOSE) $(BUILD_NATIVE_CMD) make test-native-internal

test-driver: | check-gopath $(BUILD_PATH) $(DOCKER_BUILD_NATIVE_IMAGE)
	@$(RUN_VERBOSE) $(GO_TEST) ./driver/...

build: | build-native build-driver $(DOCKER_IMAGE_VERSIONED)
build-native: $(BUILD_PATH) $(DOCKER_BUILD_NATIVE_IMAGE)
	@$(RUN) $(BUILD_NATIVE_CMD) make build-native-internal

build-driver: | check-gopath $(BUILD_PATH) $(DOCKER_BUILD_DRIVER_IMAGE) build-native
	@$(RUN) $(BUILD_DRIVER_CMD) make build-driver-internal

build-driver-internal: $(BUILD_PATH)
	@$(RUN) $(bblfsh-sdk-tools) build $(DRIVER_VERSION) --output $(BUILD_PATH)/etc/manifest.toml
	@cd driver; \
	$(RUN) $(GO_CMD) build -o $(BUILD_PATH)/bin/driver .

integration-test: build
	CONTAINER_ID=` + "`" + `$(DOCKER_CMD) run -d \
		-p $(GRPC_PORT_INTEGRATION):$(GRPC_PORT) \
		$(call unescape_docker_tag,$(DOCKER_IMAGE_VERSIONED))` + "`" + `; \
	echo "CONTAINER_ID: $$CONTAINER_ID"; \
	$(bblfsh-sdk-tools) test --endpoint localhost:$(GRPC_PORT_INTEGRATION) || true; \
	docker kill $$CONTAINER_ID;

push: build
	$(if $(pushdisabled),$(error $(pushdisabled)))

	@if [ "$$DOCKER_USERNAME" != "" ]; then \
		$(DOCKER_CMD) login -u="$$DOCKER_USERNAME" -p="$$DOCKER_PASSWORD"; \
	fi;

	@$(RUN) $(DOCKER_PUSH) $(call unescape_docker_tag,$(DOCKER_IMAGE_VERSIONED))
	@if [ "$$TRAVIS_TAG" != "" ]; then \
		$(RUN) $(DOCKER_TAG) $(call unescape_docker_tag,$(DOCKER_IMAGE_VERSIONED)) \
			$(call unescape_docker_tag,$(DOCKER_IMAGE)):latest; \
		$(RUN) $(DOCKER_PUSH) $(call unescape_docker_tag,$(DOCKER_IMAGE):latest); \
	fi;

validate:
	@$(RUN) $(bblfsh-sdk) update --dry-run

clean:
	@$(RUN) rm -rf $(BUILD_PATH)
`)

func makeRulesMkBytes() ([]byte, error) {
	return _makeRulesMk, nil
}

func makeRulesMk() (*asset, error) {
	bytes, err := makeRulesMkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "make/rules.mk", size: 4230, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"Makefile":                              makefile,
	"etc/build/Dockerfile.build.alpine.tpl": etcBuildDockerfileBuildAlpineTpl,
	"etc/build/Dockerfile.build.debian.tpl": etcBuildDockerfileBuildDebianTpl,
	"etc/run.sh":                            etcRunSh,
	"make/bootstrap.mk":                     makeBootstrapMk,
	"make/functions.mk":                     makeFunctionsMk,
	"make/help.mk":                          makeHelpMk,
	"make/rules.mk":                         makeRulesMk,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"Makefile": &bintree{makefile, map[string]*bintree{}},
	"etc": &bintree{nil, map[string]*bintree{
		"build": &bintree{nil, map[string]*bintree{
			"Dockerfile.build.alpine.tpl": &bintree{etcBuildDockerfileBuildAlpineTpl, map[string]*bintree{}},
			"Dockerfile.build.debian.tpl": &bintree{etcBuildDockerfileBuildDebianTpl, map[string]*bintree{}},
		}},
		"run.sh": &bintree{etcRunSh, map[string]*bintree{}},
	}},
	"make": &bintree{nil, map[string]*bintree{
		"bootstrap.mk": &bintree{makeBootstrapMk, map[string]*bintree{}},
		"functions.mk": &bintree{makeFunctionsMk, map[string]*bintree{}},
		"help.mk":      &bintree{makeHelpMk, map[string]*bintree{}},
		"rules.mk":     &bintree{makeRulesMk, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
