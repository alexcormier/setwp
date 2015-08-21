#!/usr/bin/env bash

if ! [[ $1 =~ [0-9\.\-] ]]; then
    echo 'invalid version number'
    exit 1
fi

VERSION="$1"
PROJECT_NAME='setwp'
PROJECT_PATH="github.com/alexandrecormier/${PROJECT_NAME}"
PROJECT_FULL_PATH="${GOPATH}/src/${PROJECT_PATH}"
BINARY_PATH="${GOPATH}/bin"
TARBALL="${PROJECT_FULL_PATH}/releases/${PROJECT_NAME}-v${VERSION}.tar.gz"
BREW_FORMULA="${PROJECT_FULL_PATH}/homebrew/${PROJECT_NAME}.rb"
BREW_FORMULA_TEMPLATE="${BREW_FORMULA}.template"
ARGS_FILE="${PROJECT_FULL_PATH}/args/args.go"

# Bump version
sed -E -i '' "s/version [0-9\.\-]+/version ${VERSION}/" "${ARGS_FILE}"

# Build
go install "${PROJECT_PATH}"

# Make tarballs
mkdir -p releases
tar -czf "${TARBALL}" completion/* -C "${BINARY_PATH}" "${PROJECT_NAME}"

# Update homebrew formula
SHA256_HASH=`openssl dgst -sha256 ${TARBALL} | xxd -ps | tr -d '\n'`
sed -e "s/<VERSION>/${VERSION}/g" \
    -e "s/<SHA256_HASH>/${SHA256_HASH}/" \
    "${BREW_FORMULA_TEMPLATE}" > "${BREW_FORMULA}"
