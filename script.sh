#!/bin/bash
# set -e
# set -x

./stackoverflow-go -test.run=StackOverflowGo
./stackoverflow-go -test.run=StackOverflowCgo
