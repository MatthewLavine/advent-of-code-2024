#!/bin/bash

# Check if a day is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

# Check if the day directory exists
day_dir="./day$1"
if [ -d "$day_dir" ]; then
    echo "Directory for day $1 already exists."
    exit 2
fi

echo "Setting up for day $1"

echo "Creating directory"

mkdir $day_dir

echo "Creating blank input files"

touch $day_dir/input.txt
touch $day_dir/demo.txt

echo "Copying template solution files"

cp template/dayX.go $day_dir/day$1.go
cp template/dayX_test.go $day_dir/day$1_test.go

echo "Done"
