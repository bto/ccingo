#!/bin/bash
TOP_DIR=$(cd $(dirname $0); pwd)
BUILD_DIR=$TOP_DIR/build
C_DIR=$TOP_DIR/c

GO_FILE=$TOP_DIR/main.go
LL_FILE=$BUILD_DIR/main.ll
AS_FILE=$BUILD_DIR/main.s
EXEC_FILE=$BUILD_DIR/main

function try() {
    expected=$1
    input=$2

    echo -n "$input" | go run $GO_FILE > $LL_FILE
    llc $LL_FILE -o=$AS_FILE
    gcc $AS_FILE $C_DIR/*.o -o $EXEC_FILE
    $EXEC_FILE
    ret=$?

    if [ $ret = $expected ]; then
        echo "OK: $input => $expected"
    else
        echo "Failed: $input => $ret, expected: $expected"
    fi
}

try 3 'main(){1 + 2;}'
try 2 'main(){  1+5
  -4  ;}'
try 7 'main(){1 + 2*3;}'
try 2 'main(){8 - 4/ 2 * 3;}'
try 9 'main(){(1 + 2)*3;}'
try 27 'main(){(1 + 2*4)*3;}'
try 21 'main(){(-1 + +2*+4)*3;}'
try 0 'main(){(1+2)==2;}'
try 1 'main(){1 ==(1 == 1);}'
try 1 'main(){(1+2) != 2;}'
try 0 'main(){3 != 1 * 3;}'
try 0 'main(){2 <  1;}'
try 0 'main(){2 <  2;}'
try 1 'main(){2 <  3;}'
try 0 'main(){2 <= 1;}'
try 1 'main(){2 <= 2;}'
try 1 'main(){2 <= 3;}'
try 1 'main(){2 >  1;}'
try 0 'main(){2 >  2;}'
try 0 'main(){2 >  3;}'
try 1 'main(){2 >= 1;}'
try 1 'main(){2 >= 2;}'
try 0 'main(){2 >= 3;}'
try 1 'main(){a = 1;}'
try 1 'main(){a = 1; b = 2; a;}'
try 2 'main(){a = 1; b = 2; b;}'
try 8 'main(){a = 1; b = 2; (a + (c = 3)) * b;}'
try 18 'main(){a = b = (1 == 1) * 3; (a + b) * b;}'
try 1 'main(){foo = 1; return foo; bar = 2; (foo + bar) * bar;}'
try 2 'main(){foo = 1; bar = 2; return bar; (foo + bar) * bar;}'
try 8 'main(){foo = 1; bar = 2; return (foo + (baz = 3)) * bar;}'
try 1 'main(){foo = 1; bar = 2; (foo + bar) * bar; return foo;}'
try 1 'main(){foo=1;if(0)foo=2;foo;}'
try 2 'main(){foo=1;if(1)foo=2;foo;}'
try 10 'main(){i=0;while(i<10)i=i+1;i;}'
try 10 'main(){for(i=0;i<10;i=i+1)i;i;}'
try 4 'main(){foo=1;bar=2;if(1){foo=3;bar=4;}bar;}'
try 255 'main(){retuchar();}'
try 1 'main(){add1(1);}'
try 3 'main(){add2(1,2);}'
try 6 'main(){add3(1,2,3);}'
try 10 'main(){add4(1,2,3,4);}'
try 15 'main(){add5(1,2,3,4,5);}'
try 21 'main(){add6(1,2,3,4,5,6);}'
try 89 'main(){fib(10);}'
