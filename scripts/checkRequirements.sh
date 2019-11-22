#!/bin/bash

function e {
    echo -e "\e[32m\e[1m### $1 ###\e[0m"
}

e "checking go environment"
VARIABLES=(GOPATH GOBIN)
for i in ${VARIABLES[@]}; do
    if [ -z ${i+x} ];
    then
        echo "$i is unset. aborting.";
        exit 1;
    fi
done

e "checking for required commands"
REQ=(
    go
    make
    statik)

for i in ${REQ[@]}; do
    command -v $i >/dev/null 2>&1 || { echo >&2 "$i is required but it's not installed.  Aborting."; exit 1; }
done

e "checking go environment complete"
