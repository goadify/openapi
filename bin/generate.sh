#!/bin/sh

generate_go() {
    cd "$1"/go && go generate
}

"$@"