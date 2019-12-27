#!/bin/bash

echo "Building API"
CGO_ENABLED=0 go build -o /docker/api main.go

echo "Creating docker image"
cd docker && docker build -t cntr-multitool .

echo "Finished!"
echo "Start container with the following command:"
echo "docker run --rm -d -p 9100:9100 cnt-multitool"

