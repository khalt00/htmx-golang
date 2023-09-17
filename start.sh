#!/bin/bash

echo "really doing sthing"

start(){
    echo "building and starting"
    go build -v -o htmx && ./htmx
}
$*
