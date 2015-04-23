#!/usr/bin/env bash

if ! [[ $1 =~ [0-9\.\-] ]]; then
    echo 'invalid version number'
    exit 1
fi

VERSION="$1"
PROJECT_NAME='setwp'
PROJECT_PATH="github.com/alexandrecormier/${PROJECT_NAME}"
PROJECT_FULL_PATH="${GOPATH}/src/${PROJECT_PATH}"
BINARY_PATH_32BIT="${GOPATH}/bin/darwin_386"
BINARY_PATH_64BIT="${GOPATH}/bin"
TARBALL_32BIT="${PROJECT_FULL_PATH}/releases/${PROJECT_NAME}-i386-v${VERSION}.tar.gz"
TARBALL_64BIT="${PROJECT_FULL_PATH}/releases/${PROJECT_NAME}-amd64-v${VERSION}.tar.gz"
BREW_FORMULA="${PROJECT_FULL_PATH}/homebrew/${PROJECT_NAME}.rb"
BREW_FORMULA_TEMPLATE="${BREW_FORMULA}.template"
ARGS_FILE="${PROJECT_FULL_PATH}/args/args.go"

# Build
go install "${PROJECT_PATH}"
GOARCH=386 CGO_ENABLED=1 go install "${PROJECT_PATH}"

# Make tarballs
mkdir -p releases
tar -czf "${TARBALL_32BIT}" completion/* -C "${BINARY_PATH_32BIT}" "${PROJECT_NAME}"
tar -czf "${TARBALL_64BIT}" completion/* -C "${BINARY_PATH_64BIT}" "${PROJECT_NAME}"

# Update homebrew formula
SHA256_32BIT=`openssl dgst -sha256 ${TARBALL_32BIT} | xxd -ps | tr -d '\n'`
SHA256_64BIT=`openssl dgst -sha256 ${TARBALL_64BIT} | xxd -ps | tr -d '\n'`
sed -e "s/<VERSION>/${VERSION}/g" \
    -e "s/<SHA256_32BIT>/${SHA256_32BIT}/" \
    -e "s/<SHA256_64BIT>/${SHA256_64BIT}/" \
    "${BREW_FORMULA_TEMPLATE}" > "${BREW_FORMULA}"

# Bump version
sed -E -i '' "s/version [0-9\.\-]+/version ${VERSION}/" "${ARGS_FILE}"
