#!/bin/bash

RED="\033[1;31m"
GREEN="\033[1;32m"
NOCOLOR="\033[0m"

# Run install script.
./install.sh

# Build from source.
go build *.go
if [ $? -eq 0 ]; then
    echo -e "${GREEN}Build done.${NOCOLOR}"
else
    echo -e "${RED}Build failed.${NOCOLOR}"
    exit 1
fi

# Run tests and check the results
c=$(go test -p 1 ./... -v)
if [[ $c = *"FAIL"* ]]; then
    echo -e "${RED}Test fails${NOCOLOR}"
    echo "$c"
    rm document_api # Remove built app
    exit 1
else
    echo -e "${GREEN}All tests passed.${NOCOLOR}"
fi

# Check for bin directory.
if [ ! -d "bin" ]; then
    # Create bin directory once its missing.
    mkdir bin/
fi

# Move App to bin folder.
mv document_api bin/
if [ $? -eq 0 ]; then
    echo "Application moved into the build folder."
else
    echo -e "${RED}Error${NOCOLOR} while moving application to /bin folder."
    exit 1
fi

# Remove debug.test files
rm -rf */debug.test

# Run application
./bin/document_api