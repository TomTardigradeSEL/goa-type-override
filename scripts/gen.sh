#! /bin/bash

if [ "$(basename "$(pwd)")" != "goa-type-override" ]; then
    print_error "Must run from repo root directory."
    print_help
    return 1
fi

rm -rf interfaces

protoc \
    -I=proto/pets/ \
    --go_out=. \
    --go_opt=module=petsvc \
    ./proto/pets/*.proto

rm -rf gen

goa gen petsvc/design
