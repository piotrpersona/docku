#!/usr/bin/env bash

PACKAGE="${PACKAGE:-"main.go"}"
PACKAGE_NAME="${PACKAGE_NAME:-"docku"}"

PLATFORMS=(
    "darwin/386"
    "darwin/amd64"
    "darwin/arm"
    "darwin/arm64"
    "dragonfly/amd64"
    "freebsd/386"
    "freebsd/amd64"
    "freebsd/arm"
    "linux/386"
    "linux/amd64"
    "linux/arm"
    "linux/arm64"
    "linux/ppc64"
    "linux/ppc64le"
    "linux/mips"
    "linux/mipsle"
    "linux/mips64"
    "linux/mips64le"
    "netbsd/386"
    "netbsd/amd64"
    "netbsd/arm"
    "openbsd/386"
    "openbsd/amd64"
    "openbsd/arm"
    "plan9/386"
    "plan9/amd64"
    "solaris/amd64"
)

for platform in "${PLATFORMS[@]}"; do
    PLATFORM_SPLIT=( ${platform//\// } )
    GOOS=${PLATFORM_SPLIT[0]}
    GOARCH="${PLATFORM_SPLIT[1]}"
    OUTPUT_NAME="${PACKAGE_NAME}-${GOOS}-${GOARCH}"
    if [[ "${GOOS}" = "windows" ]]; then
        OUTPUT_NAME+='.exe'
    fi

    env GOOS="${GOOS}" GOARCH="${GOARCH}" go build -o "${OUTPUT_NAME}" "${PACKAGE}"
    if [[ "${?}" -ne 0 ]]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
