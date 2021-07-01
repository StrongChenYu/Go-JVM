#!/bin/sh
# set -ex

# cd v1/code/go
# export GOPATH=$PWD

# export JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk1.8.0_144.jdk/Contents/Home/
go run go-jvm/ch01 -version | grep -q "version 0.0.1"
echo "pass test in ch01"

go run go-jvm/ch02 java.lang.Object | grep -q "class data"
echo "pass test in ch02"

go run go-jvm/ch03 java.lang.Object | grep -q "this class: java/lang/Object"
echo "pass test in ch03"

# go run go-jvm/ch04 java.lang.Object 2>&1 | grep -q "100"
# go run go-jvm/ch05 -cp ../example.jar go-jvm.book.ch05.GaussTest 2>&1 | grep -q "5050"
# go run go-jvm/ch06 -cp ../example.jar go-jvm.book.ch06.MyObject | grep -q "32768"
# go run go-jvm/ch07 -cp ../example.jar go-jvm.book.ch07.FibonacciTest | grep -q "832040"
# go run go-jvm/ch08 -cp ../example.jar go-jvm.book.ch01.HelloWorld  | grep -q "Hello, world!"
# go run go-jvm/ch08 -cp ../example.jar go-jvm.book.ch08.PrintArgs foo bar | tr -d '\n' | grep -q "foobar"
# go run go-jvm/ch09 -cp ../example.jar go-jvm.book.ch09.GetClassTest | grep -q "Ljava.lang.String;"
# go run go-jvm/ch09 -cp ../example.jar go-jvm.book.ch09.StringTest | tr -d '\n' | grep -q "truefalsetrue"
# go run go-jvm/ch09 -cp ../example.jar go-jvm.book.ch09.ObjectTest | tr -d '\n' | grep -q "falsetrue"
# go run go-jvm/ch09 -cp ../example.jar go-jvm.book.ch09.CloneTest | grep -q "3.14"
# go run go-jvm/ch09 -cp ../example.jar go-jvm.book.ch09.BoxTest | grep -q "1, 2, 3"
# go run go-jvm/ch10 -cp ../example.jar go-jvm.book.ch10.ParseIntTest 123 | grep -q "123"
# go run go-jvm/ch10 -cp ../example.jar go-jvm.book.ch10.ParseIntTest abc 2>&1 | grep  'For input string: "abc"'
# go run go-jvm/ch10 -cp ../example.jar go-jvm.book.ch10.ParseIntTest 2>&1 | grep -q "at go-jvm"
# go run go-jvm/ch11 -cp ../example.jar go-jvm.book.ch01.HelloWorld  | grep -q "Hello, world!"

echo OK