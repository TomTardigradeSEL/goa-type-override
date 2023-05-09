#! /bin/bash

if [ "$(basename "$(pwd)")" != "goa-type-override" ]; then
    print_error "Must run from repo root directory."
    print_help
    return 1
fi

go install github.com/bwplotka/bingo@latest
bingo get -l github.com/bwplotka/bingo
bingo get -l goa.design/goa/v3/cmd/goa@v3.11.3
bingo get -l google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
bingo get -l google.golang.org/protobuf/cmd/protoc-gen-go@v1.29.0
