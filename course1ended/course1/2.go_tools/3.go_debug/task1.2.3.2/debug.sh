#!/bin/bash

echo "Debug started..."
go build -o myprogram main.go
dlv exec ./myprogram
echo "Debug ended."