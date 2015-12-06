#!/bin/bash

for f in *.go
do
	echo Building $f
	go build -ldflags "-s" $f
	
done
