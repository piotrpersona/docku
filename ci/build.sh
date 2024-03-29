#!/usr/bin/env bash

function __help() {
    echo "build docku binary"
    echo
    echo "Supported platforms: ${PLATFORMS[@]}"
    echo
    echo "Usage"
    echo
    echo "./build.sh [PLATFORMS...]"
    echo
    echo "Options:"
    echo
    echo "  -h --help    Display help message"
}

function __config() {
    PACKAGE="${PACKAGE:-"main.go"}"
    PACKAGE_NAME="${PACKAGE_NAME:-"docku"}"

    PLATFORMS=(
        "darwin/386"
        "darwin/amd64"
        "linux/386"
        "linux/amd64"
    )

    GIT_TAG="$( git describe --abbrev=0 )"
    GIT_HASH="$( git rev-parse HEAD )"
    BUILD_DATE="$( date +%F )"
    DOCKU_REPO="github.com/piotrpersona/docku"
    read -r -d '' LDFLAGS << EOM
        -X ${DOCKU_REPO}/cmd.gitVersionTag=${GIT_TAG}
        -X ${DOCKU_REPO}/cmd.gitHash=${GIT_HASH}
        -X ${DOCKU_REPO}/cmd.buildDate=${BUILD_DATE}
EOM
}

function __handle_args() {
    declare -a CUSTOM_PLATFORMS
    while [[ "${#}" -gt 0 ]]; do
        case "${1}" in
            -h|--help) __help; exit 0;;
            -*) echo "Unkown options: ${1}"; exit 1;;
            *) CUSTOM_PLATFORMS+=( "${1}" ); shift 1;
        esac
    done

    [[ -z "${CUSTOM_PLATFORMS}" ]] \
        || PLATFORMS=( "${CUSTOM_PLATFORMS[@]}" )
}

function __build_docku() {
    for platform in "${PLATFORMS[@]}"; do
        local PLATFORM_SPLIT=( ${platform//\// } )
        local GOOS=${PLATFORM_SPLIT[0]}
        local GOARCH="${PLATFORM_SPLIT[1]}"
        local output_name="${OUTPUT_NAME:-"${PACKAGE_NAME}-${GOOS}-${GOARCH}"}"
        if [[ "${GOOS}" = "windows" ]]; then
            output_name+='.exe'
        fi

        env GOOS="${GOOS}" GOARCH="${GOARCH}" go build -ldflags "${LDFLAGS}" \
            -o "${GOPATH}/bin/${output_name}" "${PACKAGE}"
        if [[ "${?}" -ne 0 ]]; then
            echo 'An error has occurred! Aborting the script execution...'
            exit 1
        fi
    done
}

function main() {
    __config
    __handle_args "${@}"
    __build_docku
}

main "${@}"
