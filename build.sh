#!/usr/bin/env bash

PACKAGE="${PACKAGE:-"main.go"}"
PACKAGE_NAME="${PACKAGE_NAME:-"docku"}"

PLATFORMS=(
    "darwin/386"
    "darwin/amd64"
    "linux/386"
    "linux/amd64"
)

for platform in "${PLATFORMS[@]}"; do
    PLATFORM_SPLIT=( ${platform//\// } )
    GOOS=${PLATFORM_SPLIT[0]}
    GOARCH="${PLATFORM_SPLIT[1]}"
    OUTPUT_NAME="${PACKAGE_NAME}-${GOOS}-${GOARCH}"
    if [[ "${GOOS}" = "windows" ]]; then
        OUTPUT_NAME+='.exe'
    fi

    env GOOS="${GOOS}" GOARCH="${GOARCH}" go build -o "${GOPATH}/bin/${OUTPUT_NAME}" "${PACKAGE}"
    if [[ "${?}" -ne 0 ]]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
