#!/bin/bash
SRC_DIR=$(cd $(dirname $0); pwd)
TOP_DIR=$(cd $SRC_DIR/../..; pwd)
source $SRC_DIR/../common.sh

try 3 '1 + 2'
try 2 '  1+5
  -4  '
