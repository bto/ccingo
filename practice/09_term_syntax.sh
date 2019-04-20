#!/bin/bash
source $(dirname $0)/common.sh

try 3 '1 + 2'
try 2 '  1+5
  -4  '
try 7 '1 + 2*3'
try 2 '8 - 4/ 2 * 3'
try 9 '(1 + 2)*3'
try 27 '(1 + 2*4)*3'
