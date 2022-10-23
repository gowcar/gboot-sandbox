GO := go
#GOOS := linux
GOOS := darwin
GO_ENVS := CGO_ENABLED=0 GOOS=${GOOS} GOPROXY=direct GOSUMDB=off

BUILD_TIME = $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_SHA := $(shell git rev-parse HEAD 2>/dev/null)
VERSION := 0.0.1
ENV := Release

GO_LD_FLAGS = \
	-X main.BuildTime=${BUILD_TIME} \
	-X main.GitSHA=${GIT_SHA} \
	-X main.Version=${VERSION} \
	-X main.Env=${ENV}

GO_BUILD_FLAGS = -v -ldflags='${GO_LD_FLAGS}'

BINARIES = gboot-demo
VERSION_NAME = ${BINARIES}-${VERSION}
DIST := dist/${VERSION_NAME}
TAR = ${VERSION_NAME}.tar.gz
PKGS = $(shell ${GO} list ./... | tr '\n' ',')
EXEC := ${DIST}/bin/${BINARIES}
#EXEC := ./${BINARIES}
CMD := ./

.PHONY: all gen build clean pkg deploy fmt check-fmt upload-swagger

all: clean pkg

gen:
	${GO} generate ./...

build:
	env ${GO_ENVS} ${GO} build ${GO_BUILD_FLAGS} -o ${EXEC} ${CMD}

clean:
	rm -rf ${TAR}
	rm -rf dist

pkg: build
	chmod a+x ${EXEC}
	mkdir -p ${DIST}/config
	cp -f config/config.yml ${DIST}/config/config.yml
	tar -zcf ${TAR} -C dist ${VERSION_NAME}
	@echo "Build successfully"