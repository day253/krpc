#!/bin/bash
set -xv
SCRIPT_PATH="$(
    cd $(dirname "$0")
    pwd
)"
TOP_DIR="${SCRIPT_PATH%/*}"
BIN_DIR="${TOP_DIR}/bin"

PROC_NAME="boilerplate"

cd ${SCRIPT_PATH} || exit 1

exec ${BIN_DIR}/${PROC_NAME}
