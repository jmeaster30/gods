#!/bin/bash
go test -v -coverprofile="coverage.text" ./linear ./composite
go tool cover -html="coverage.text" -o "coverage.html"

while getopts o: arg;
do
  case $arg in
    o) 
      firefox coverage.html
      ;;
  esac
done
