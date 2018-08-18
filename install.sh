#!/bin/bash

RED="\033[1;31m"
GREEN="\033[1;32m"
NOCOLOR="\033[0m"

# Clean go packages
# go clean

# Install go packages
# go install *.go
# if [ $? -eq 0 ]; then
#     echo -e "All packages installed"
# else
#     echo -e "An ${RED}error${NOCOLOR} occured"
#     exit 1
# fi