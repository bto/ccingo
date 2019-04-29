#!/bin/bash
SRC_DIR=$(cd $(dirname $0); pwd)
TOP_DIR=$(cd $SRC_DIR/../..; pwd)
source $SRC_DIR/../common.sh

try 3 '1 + 2'
try 2 '  1+5
  -4  '
try 7 '1 + 2*3'
try 2 '8 - 4/ 2 * 3'
try 9 '(1 + 2)*3'
try 27 '(1 + 2*4)*3'
try 21 '(-1 + +2*+4)*3'
try 0 '(1+2)==2'
try 1 '1 ==(1 * 1)'
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
