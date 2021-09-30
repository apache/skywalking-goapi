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
  find "$BASEDIR"/satellite -name "*.go" -exec rm {} \;
}

function prepareSatelliteProtocols() {
  mkdir -p "$PROTOCOLDIR/satellite/"
  cp -R "$BASEDIR"/satellite/data/v1/*.proto "$PROTOCOLDIR/satellite/"
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
    --proto_path="$PROTOCOLDIR"/envoy/api \
    --proto_path="$PROTOCOLDIR"/udpa \
    --proto_path="$PROTOCOLDIR"/protoc-gen-validate \
    --proto_path="$PROTOCOLDIR"/prometheus-model \
    --proto_path="$PROTOCOLDIR" \
    --go_opt=Mudpa/annotations/migrate.proto=github.com/cncf/xds/go/udpa/annotations \
    --go_opt=Mudpa/annotations/status.proto=github.com/cncf/xds/go/udpa/annotations \
    --go_opt=Mudpa/annotations/versioning.proto=github.com/cncf/xds/go/udpa/annotations \
    --go_opt=Menvoy/api/v2/core/socket_option.proto=github.com/envoyproxy/go-control-plane/envoy/api/v2/core \
    --go_opt=Menvoy/config/core/v3/socket_option.proto=github.com/envoyproxy/go-control-plane/envoy/config/core/v3 \
    --go_opt=Menvoy/config/core/v3/address.proto=github.com/envoyproxy/go-control-plane/envoy/config/core/v3 \
    --go_opt=Menvoy/config/core/v3/backoff.proto=github.com/envoyproxy/go-control-plane/envoy/config/core/v3 \
    --go_opt=Menvoy/config/core/v3/http_uri.proto=github.com/envoyproxy/go-control-plane/config/core/v3 \
    --go_opt=Menvoy/api/v2/core/address.proto=github.com/envoyproxy/go-control-plane/envoy/api/v2/core \
    --go_opt=Menvoy/api/v2/core/backoff.proto=github.com/envoyproxy/go-control-plane/envoy/api/v2/core \
    --go_opt=Menvoy/api/v2/core/http_uri.proto=github.com/envoyproxy/go-control-plane/envoy/api/v2/core \
    --go_opt=Menvoy/type/percent.proto=github.com/envoyproxy/go-control-plane/envoy/type \
    --go_opt=Menvoy/type/semantic_version.proto=github.com/envoyproxy/go-control-plane/envoy/type \
    --go_opt=Menvoy/api/v2/core/base.proto=github.com/envoyproxy/go-control-plane/envoy/api/v2/core \
    --go_opt=Menvoy/data/accesslog/v2/accesslog.proto=github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2 \
    --go_opt=Menvoy/data/accesslog/v3/accesslog.proto=github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3 \
    --go_opt=Menvoy/service/accesslog/v2/als.proto=github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2 \
    --go_opt=Menvoy/service/accesslog/v3/als.proto=github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3 \
    --go_opt=Menvoy/type/v3/percent.proto=github.com/envoyproxy/go-control-plane/envoy/type/v3 \
    --go_opt=Menvoy/type/v3/semantic_version.proto=github.com/envoyproxy/go-control-plane/envoy/type/v3 \
    --go_opt=Menvoy/annotations/deprecation.proto=github.com/envoyproxy/go-control-plane/envoy/annotations \
    --go_opt=Menvoy/config/core/v3/base.proto=github.com/envoyproxy/go-control-plane/envoy/config/core/v3 \
    --go_opt=Mxds/core/v3/context_params.proto=github.com/cncf/udpa/go/xds/core/v3 \
    --go_opt=Menvoy/service/accesslog/v2/als.proto=github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2 \
    --go_opt=Menvoy/service/accesslog/v3/als.proto=github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3 \
    --go_opt=Menvoy/service/metrics/v2/metrics_service.proto=github.com/envoyproxy/go-control-plane/envoy/service/metrics/v2 \
    --go_opt=Menvoy/service/metrics/v3/metrics_service.proto=github.com/envoyproxy/go-control-plane/envoy/service/metrics/v3 \
    --go_out="$BASEDIR" \
    "$PROTOCOLDIR"/satellite/*.proto

  mv "$BASEDIR"/skywalking.apache.org/repo/goapi/collect "$BASEDIR"/ \
  && mv "$BASEDIR"/skywalking.apache.org/repo/goapi/satellite/data/v1/* "$BASEDIR"/satellite/data/v1 \
  && rm -rf "$BASEDIR"/skywalking.apache.org
  go mod tidy
}

initProtocolHome
addProtocol skywalking-collect https://github.com/apache/skywalking-data-collect-protocol/archive/"${COLLECT_PROTOCOL_SHA}".tar.gz
addProtocol envoy https://github.com/envoyproxy/envoy/archive/"${ENVOY_SERVICE_PROTOCOL_SHA}".tar.gz
addProtocol udpa https://github.com/cncf/udpa/archive/"${UDPA_SERVICE_PROTOCOL_SHA}".tar.gz
addProtocol protoc-gen-validate https://github.com/envoyproxy/protoc-gen-validate/archive/${ENVOY_PROTOC_VALIDATE_SHA}.tar.gz
addProtocol prometheus-model https://github.com/prometheus/client_model/archive/${PROMETHEUS_MODEL_SHA}.tar.gz
cleanHistoryCodes
prepareSatelliteProtocols
generateCodes