# Build on macOS

## Build PuzzleDB

Go to the directory where you want to download the Vitess source code and clone the Vitess GitHub repo:

    git clone https://github.com/cybergarage/puzzledb-go
    cd puzzledb-go

To build PuzzleDB, run the following command:

    make build

## Testing PuzzleDB

PuzzleDB uses the Go testing framework. To run all tests, run the following command:

    make test
