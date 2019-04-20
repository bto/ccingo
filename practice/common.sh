PRACTICE_DIR=$(dirname $0)
TOP_DIR=$PRACTICE_DIR/..
BUILD_DIR=$TOP_DIR/build

GO_BASENAME=$(basename $0)
GO_BASENAME=${GO_BASENAME%.sh}
GO_FILE=$PRACTICE_DIR/$GO_BASENAME.go
AS_FILE=$BUILD_DIR/$GO_BASENAME.s
EXEC_FILE=$BUILD_DIR/$GO_BASENAME

function try() {
    expected=$1
    input=$2

    echo -n $input | go run $GO_FILE > $AS_FILE
    gcc $AS_FILE -o $EXEC_FILE
    $EXEC_FILE
    ret=$?

    if [ $ret = $expected ]; then
        echo "OK: $input => $expected"
    else
        echo "Failed: $input => $expected"
    fi
}
