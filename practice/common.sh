BUILD_DIR=$TOP_DIR/build
C_DIR=$TOP_DIR/practice/c
TARGET_DIR=$BUILD_DIR/${SRC_DIR#$TOP_DIR/}

GO_FILE=$SRC_DIR/main.go
AS_FILE=$TARGET_DIR/main.s
EXEC_FILE=$TARGET_DIR/main

mkdir -p $TARGET_DIR

function try() {
    expected=$1
    input=$2

    echo -n "$input" | go run $GO_FILE > $AS_FILE
    gcc $AS_FILE $C_DIR/*.o -o $EXEC_FILE
    $EXEC_FILE
    ret=$?

    if [ $ret = $expected ]; then
        echo "OK: $input => $expected"
    else
        echo "Failed: $input => $ret, expected: $expected"
    fi
}
