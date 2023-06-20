#!/bin/sh

UID=$(id -u)

generate_go() {
    cd "$1"/go && go generate
}

generate_ts() {
  cd "$1" && docker run --rm \
    -u "${UID}":"${UID}" \
    -v "${PWD}":/local openapitools/openapi-generator-cli generate \
    -i /local/openapi.yaml \
    -g typescript-axios \
    -o /local/ts

  rm -rf ./ts/.openapi-generator ./ts/.npmignore ./ts/.openapi-generator-ignore ./ts/git_push.sh
}

"$@"