# Run install script.
./install.sh

# Build from source.
go build *.go

echo "Build done"

# Check for bin directory.
if [ ! -d "$bin" ]; then
    # Create bin directory once its missing.
    mkdir bin/
fi

# Move App to bin folder.
mv app bin/

echo "Application moved into the build folder"