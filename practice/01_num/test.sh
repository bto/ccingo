#!/bin/bash
SRC_DIR=$(cd $(dirname $0); pwd)
TOP_DIR=$(cd $SRC_DIR/../..; pwd)
source $SRC_DIR/../common.sh

try 1 1
try 2 2
