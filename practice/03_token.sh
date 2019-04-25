#!/bin/bash
source $(dirname $0)/common.sh

try 3 '1 + 2'
try 2 '  1+5
  -4  '
