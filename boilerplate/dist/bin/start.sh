#!/bin/bash
SCRIPT_PATH="$(
    cd $(dirname "$0")
    pwd
)"
TOP_DIR="${SCRIPT_PATH%/*}"
BIN_DIR="${TOP_DIR}/bin"
cd ${SCRIPT_PATH} || exit 1
${SCRIPT_PATH}/control restart
