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
export PROTOCOLDIR="$BASEDIR"/temp/protocols
. "$BASEDIR"/dependencies.sh

function initProtocolHome(){
  if [[ -d "$PROTOCOLDIR" ]]; then
    rm -rf "$PROTOCOLDIR}"
  fi
  mkdir -p "$PROTOCOLDIR"
}

function addProtocol(){
  name=$1
  protocolTarAddress=$2

  curl -sLo "$PROTOCOLDIR"/$name.tgz $2

  if [[ ! -d "$PROTOCOLDIR"/$name ]]; then
    mkdir "$PROTOCOLDIR"/$name
  else
    rm -rf "$PROTOCOLDIR"/$name/*
  fi

  tar -zxf "$PROTOCOLDIR"/$name.tgz -C "$PROTOCOLDIR"/$name --strip 1

  find "$PROTOCOLDIR"/$name -regex ".*[cC]ompat.proto" -exec rm {} \;
}

function cleanHistoryCodes(){
  rm -rf "$BASEDIR"/collect
  rm -rf "$BASEDIR"/proto
  find "$BASEDIR"/satellite -name "*.go" -exec rm {} \;
}

function prepareSatelliteProtocols() {
  mkdir -p "$PROTOCOLDIR/satellite/"
  cp -R "$BASEDIR"/satellite/data/v1/*.proto "$PROTOCOLDIR/satellite/"
  cp -R "$BASEDIR"/satellite/envoy/accesslog/v3/*.proto "$PROTOCOLDIR/satellite/"
}

function generateCodes(){
  go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0
  go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

  "$BASEDIR"/scripts/protoc.sh \
    --proto_path="$PROTOCOLDIR"/skywalking-collect \
    --go_out="$BASEDIR" \
    --go-grpc_out="$BASEDIR" \
    "$PROTOCOLDIR"/skywalking-collect/*/*.proto

  "$BASEDIR"/scripts/protoc.sh \
    --proto_path="$PROTOCOLDIR"/skywalking-collect \
    --proto_path="$PROTOCOLDIR"/envoy \
    --proto_path="$PROTOCOLDIR"/xds \
    --proto_path="$PROTOCOLDIR"/protoc-gen-validate \
    --proto_path="$PROTOCOLDIR"/prometheus-model \
    --proto_path="$PROTOCOLDIR" \
    --go_out="$BASEDIR" \
    --go-grpc_out="$BASEDIR" \
    $(bash "$BASEDIR"/scripts/envoy-import.sh opts "$PROTOCOLDIR") \
    "$PROTOCOLDIR"/satellite/*.proto \
    $(bash "$BASEDIR"/scripts/envoy-import.sh files "$PROTOCOLDIR") \

  mv "$BASEDIR"/skywalking.apache.org/repo/goapi/collect "$BASEDIR"/ \
  && mv "$BASEDIR"/skywalking.apache.org/repo/goapi/satellite/data/v1/* "$BASEDIR"/satellite/data/v1 \
  && mv "$BASEDIR"/skywalking.apache.org/repo/goapi/satellite/envoy/accesslog/v3/* "$BASEDIR"/satellite/envoy/accesslog/v3 \
  && mv "$BASEDIR"/skywalking.apache.org/repo/goapi/proto/ "$BASEDIR"/ \
  && rm -rf "$BASEDIR"/skywalking.apache.org && rm -rf $TEMPDIR

  go mod tidy
}

initProtocolHome
addProtocol skywalking-collect https://github.com/apache/skywalking-data-collect-protocol/archive/"${COLLECT_PROTOCOL_SHA}".tar.gz
addProtocol envoy https://github.com/envoyproxy/data-plane-api/archive/${ENVOY_SERVICE_PROTOCOL_SHA}.tar.gz
addProtocol xds https://github.com/cncf/xds/archive/${XDS_SERVICE_PROTOCOL_SHA}.tar.gz
addProtocol protoc-gen-validate https://github.com/envoyproxy/protoc-gen-validate/archive/${PROTOC_VALIDATE_SHA}.tar.gz
addProtocol prometheus-model https://github.com/prometheus/client_model/archive/${PROMETHEUS_MODEL_SHA}.tar.gz
cleanHistoryCodes
prepareSatelliteProtocols
generateCodes