#!/bin/bash

function e {
    echo -e "\e[32m\e[1m### $1 ###\e[0m"
}

OUT=../generated

e "running protobuf code generation"
cd api
for i in *.proto; do
    echo "processing $i"
    basename="${i%.*}"
    mkdir -p ../generated/$basename
    mkdir -p ../generated/$basename/clients/java
    mkdir -p ../generated/$basename/clients/python
    OUT=../generated/$basename
    protoc -I . -I=${GOPATH}/src -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
        -I=${GOPATH}/src/github.com/gogo/googleapis/ \
        -I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ \
        $i \
        --gogoslick_out=plugins=grpc:$OUT \
        --grpc-gateway_out=logtostderr=true:$OUT \
        --swagger_out=logtostderr=true:$OUT \
        --doc_out=markdown,${basename}.md:../docs \
        --java_out=$OUT/clients/java \
        --python_out=$OUT/clients/python
    cp $OUT/*.swagger.json ../web/swagger-ui/swaggers
done

cd ../web/swagger-ui
urls="\t\turls: [{"
for f in swaggers/*.swagger.json; do
    echo $f
    urls+="{url: \"$f\", name: \"$f\"},"
done
urls="${urls::-1}"
urls+="}]"
echo "updating swagger-ui with $urls"

sed -i "s|\t\turls.*|$urls|g" index.html

e "protobuf code generation complete"