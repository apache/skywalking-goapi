#!/usr/bin/env bash

# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

set -e


export BASEDIR=$(dirname "$0")/..
export TEMPDIR="$BASEDIR"/temp
. "$BASEDIR"/dependencies.sh

function makeProtocolHome(){
    if [[ ! -d "$TEMPDIR" ]]; then
      mkdir -p "$TEMPDIR"
    else
      rm -rf "${TEMPDIR:?}"/*
    fi

    curl -sLo "$TEMPDIR"/collect-protocol.tgz https://github.com/apache/skywalking-data-collect-protocol/archive/"${COLLECT_PROTOCOL_SHA}".tar.gz

    if [[ ! -d "$TEMPDIR"/collect-protocol ]]; then
      mkdir "$TEMPDIR"/collect-protocol
    else
      rm -rf "$TEMPDIR"/collect-protocol/*
    fi

    tar -zxf "$TEMPDIR"/collect-protocol.tgz -C "$TEMPDIR"/collect-protocol --strip 1

    find "$TEMPDIR"/collect-protocol -name "*Compat.proto" -exec rm {} \;

    if [[ ! -d "$TEMPDIR"/collect-protocol/satellite ]]; then
      mkdir "$TEMPDIR"/collect-protocol/satellite
    else
      rm -rf "$TEMPDIR"/collect-protocol/satellite/*
    fi

    cp -R "$BASEDIR"/satellite/data/v1/*.proto "$TEMPDIR"/collect-protocol/satellite/
}


function cleanHistoryCodes(){
    rm -rf "$BASEDIR"/collect
    find "$BASEDIR"/satellite -name "*.go" -exec rm {} \;
}


function generateCodes(){
    go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

    "$BASEDIR"/scripts/protoc.sh \
      --proto_path="$TEMPDIR"/collect-protocol \
      --go_out="$BASEDIR" \
      --go-grpc_out="$BASEDIR" \
      "$TEMPDIR"/collect-protocol/*/*.proto

    mv "$BASEDIR"/skywalking.apache.org/repo/goapi/collect "$BASEDIR"/ \
    && mv "$BASEDIR"/skywalking.apache.org/repo/goapi/satellite/data/v1/* "$BASEDIR"/satellite/data/v1 \
    && rm -rf "$BASEDIR"/skywalking.apache.org
    go mod tidy
}

makeProtocolHome
cleanHistoryCodes
generateCodes
