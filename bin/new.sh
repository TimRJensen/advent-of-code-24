#!/bin/bash

# usage
if [ -z "$1" ]; then
  echo "Usage: $0 <day-number>"
  exit 1
fi

DIR="day$1"
mkdir -p "$DIR"

GO_FILE="$DIR/day$1.go"
TEST_FILE="$DIR/day$1_test.go"
README_FILE="$DIR/README.md"
touch "$GO_FILE"
touch "$TEST_FILE"
touch "$README_FILE"

cat <<EOL > "$GO_FILE"
package main

import (
	"log"
	"os"
)

func parse(path string) {
	buff, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
}
EOL

cat <<EOL > "$TEST_FILE"
package main
EOL

cat <<EOL > "$README_FILE"
# Solution for Advent of Code 24 day $1 

## Task 1

## Task 2

## Deliverable
EOL

# All is well
echo "- $GO_FILE"
echo "- $TEST_FILE"
echo "- $README_FILE"
