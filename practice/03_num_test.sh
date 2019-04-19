#!/bin/bash
PRACTICE_DIR=$(dirname $0)
TOP_DIR=$PRACTICE_DIR/..
BUILD_DIR=$TOP_DIR/build

GO_FILE=$PRACTICE_DIR/03_num.go
AS_FILE=$BUILD_DIR/03_num.s
EXEC_FILE=$BUILD_DIR/03_num

function try() {
    expected=$1
    input=$2

    echo $input | go run $GO_FILE > $AS_FILE
    gcc $AS_FILE -o $EXEC_FILE
    $EXEC_FILE
    ret=$?

    if [ $ret = $expected ]; then
        echo "OK: $input => $expected"
    else
        echo "Failed: $input => $expected"
    fi
}

try 1 1
try 2 2