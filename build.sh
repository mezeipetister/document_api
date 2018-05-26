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

# Check for bin directory.
if [ ! -d "bin" ]; then
    # Create bin directory once its missing.
    mkdir bin/
fi

# Move App to bin folder.
mv app bin/
if [ $? -eq 0 ]; then
    echo "Application moved into the build folder."
else
    echo -e "${RED}Error${NOCOLOR} while moving application to /bin folder."
    exit 1
fi