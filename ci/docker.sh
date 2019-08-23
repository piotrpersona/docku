#!/usr/bin/env bash

function __config() {
    SCRIPT_PATH="$( dirname "${0}" )"
    REPO="$( dirname "${SCRIPT_PATH}" )"
    DOCKERFILE="${REPO}/Dockerfile"
    DOCKER_CONTEXT="${REPO}"

    GIT_TAG="$( git describe --abbrev=0 )"
    GIT_HASH="$( git rev-parse HEAD )"

    DOCKER_REPO="${DOCKER_REPO:-"piotrpersona"}"
    IMAGE_NAME="${IMAGE_NAME:-"docku"}"
    DOCKU_IMAGE="${DOCKER_REPO}/${IMAGE_NAME}"
}

function __build() {
    OUTPUT_NAME="docku"
    TAG="${GIT_HASH}"

    docker build \
        -f "${DOCKERFILE}" \
        -t "${DOCKU_IMAGE}:${GIT_HASH}" \
        "${DOCKER_CONTEXT}"
}

# TODO: Use docku :D
function __upload() {
    TAGS=( "${GIT_TAG}" "latest" )
    for tag in "${TAGS[@]}"; do
        echo "${tag}"
        docker tag "${DOCKU_IMAGE}:${GIT_HASH}" "${DOCKU_IMAGE}:${tag}"
        docker push "${DOCKU_IMAGE}:${tag}"
    done
}

function main() {
    __config
    __build
    __upload
}

main "${@}"
