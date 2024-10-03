#!/usr/bin/env bash

# Generate a new day
mkdir -p day$1/input
cd day$1
touch input/input.txt
touch input/test.txt
echo "package main" > main.go
echo "package main" > main_test.go
