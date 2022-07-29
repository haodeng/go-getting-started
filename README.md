# go-getting-started


Examples from https://go.dev/doc/

In hello:

   # use local package
   go mod edit -replace example.com/greetings=../greetings


To build:

    cd hello
    go build
    ./hello

To install:

    go install # install to go bin dir locally