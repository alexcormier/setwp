#!/usr/bin/env bash

if ! [[ $1 =~ [0-9\.\-] ]]; then
    echo 'invalid version number'
    exit 1
fi

VERSION="$1"
SHORT_VERSION=`echo ${VERSION} | cut -f1 -d"-"`
PROJECT_NAME='setwp'
PROJECT_PATH="github.com/alexcormier/${PROJECT_NAME}"
PROJECT_FULL_PATH="${GOPATH}/src/${PROJECT_PATH}"
BINARY_PATH="${GOPATH}/bin"
TARBALL="${PROJECT_FULL_PATH}/releases/${PROJECT_NAME}-v${VERSION}.tar.gz"
ARGS_FILE="${PROJECT_FULL_PATH}/args/args.go"

# Bump version
sed -E -i '' "s/version [0-9\.]+/version ${SHORT_VERSION}/" "${ARGS_FILE}"

# Build
go install "${PROJECT_PATH}"

# Make tarball
mkdir -p releases
tar -czf "${TARBALL}" completion/* -C "${BINARY_PATH}" "${PROJECT_NAME}"
