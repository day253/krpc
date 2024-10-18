#!/bin/bash
SCRIPT_PATH="$(
    cd $(dirname "$0")
    pwd
)"
TOP_DIR="${SCRIPT_PATH%/*}"
cd ${TOP_DIR} || exit 1
cd ..
make build-in-docker
