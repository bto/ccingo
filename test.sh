#!/bin/bash
TOP_DIR=$(cd $(dirname $0); pwd)
BUILD_DIR=$TOP_DIR/build

GO_FILE=$TOP_DIR/main.go
LL_FILE=$BUILD_DIR/main.ll
AS_FILE=$BUILD_DIR/main.s
EXEC_FILE=$BUILD_DIR/main

function try() {
    expected=$1
    input=$2

    echo -n "$input" | go run $GO_FILE > $LL_FILE
    llc $LL_FILE -o=$AS_FILE
    gcc $AS_FILE -o $EXEC_FILE
    $EXEC_FILE
    ret=$?

    if [ $ret = $expected ]; then
        echo "OK: $input => $expected"
    else
        echo "Failed: $input => $ret, expected: $expected"
    fi
}

try 3 '1 + 2'
try 2 '  1+5
  -4  '
try 7 '1 + 2*3'
try 2 '8 - 4/ 2 * 3'
try 9 '(1 + 2)*3'
try 27 '(1 + 2*4)*3'
try 21 '(-1 + +2*+4)*3'
try 0 '(1+2)==2'
try 1 '1 ==(1 == 1)'
try 1 '(1+2) != 2'
try 0 '3 != 1 * 3'
try 0 '2 <  1'
try 0 '2 <  2'
try 1 '2 <  3'
try 0 '2 <= 1'
try 1 '2 <= 2'
try 1 '2 <= 3'
try 1 '2 >  1'
try 0 '2 >  2'
try 0 '2 >  3'
try 1 '2 >= 1'
try 1 '2 >= 2'
try 0 '2 >= 3'
