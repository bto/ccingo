#!/bin/bash
SRC_DIR=$(cd $(dirname $0); pwd)
TOP_DIR=$(cd $SRC_DIR/../..; pwd)
source $SRC_DIR/../common.sh

try 3 '1 + 2;'
try 2 '  1+5
  -4 ; '
try 7 '1 + 2*3 ;'
try 2 '8 - 4/ 2 * 3 ;'
try 9 '(1 + 2)*3 ;'
try 27 '(1 + 2*4)*3 ;'
try 1 'a = 1;'
try 1 'a = 1; b = 2; a;'
try 2 'a = 1; b = 2; b;'
try 8 'a = 1; b = 2; (a + (c = 3)) * b;'
try 18 'a = b = 3; (a + b) * b;'
