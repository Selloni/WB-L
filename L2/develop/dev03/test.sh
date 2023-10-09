#!/bin/bash

rm -rf dev03
go build

./dev03 text.txt > devSort.txt
sort text.txt > sort.txt
diff -s devSort.txt sort.txt
rm sort.txt devSort.txt

./dev03 -r text.txt > devSort.txt
sort -r text.txt > sort.txt
diff -s devSort.txt sort.txt
rm sort.txt devSort.txt

./dev03 -n text.txt > devSort.txt
sort -n text.txt > sort.txt
diff -s devSort.txt sort.txt
rm sort.txt devSort.txt


