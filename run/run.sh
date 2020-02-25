#!/usr/bin/env bash

function ping() {
  echo $1
}

function run_generate() {
  java -jar etc/antlr-4.8-complete.jar etc/GoLexer.g4 -Dlanguage=Go
  java -jar etc/antlr-4.8-complete.jar etc/GoParser.g4 -Dlanguage=Go
}

# https://www.tutorialspoint.com/unix/case-esac-statement
case $1 in
ping)
  ping "${@:2}"
  ;;
generate)
  run_generate
  ;;
*)
  echo "./docker/run.sh [ping|generate]"
  ;;
esac
