#!/bin/bash
source $(dirname $0)/common.sh

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
try 6 'a = 1; b = 2; (a + b) * b;'
try 18 'a = b = 3; (a + b) * b;'
try 1 'a = 1; return a; b = 2; (a + b) * b;'
try 2 'a = 1; b = 2; return b; (a + b) * b;'
try 6 'a = 1; b = 2; return (a + b) * b;'
try 1 'a = 1; b = 2; (a + b) * b; return a;'
try 6 'foo = 1; bar = 2; return (foo + bar) * bar;'
