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

TYPE=$1
PROTODIR=$2

function proto(){
  name=$1
  path=$2

  pack=$(dirname $path)
  if [[ $TYPE == "opts" ]];
  then
    echo "--go_opt=M$path=skywalking.apache.org/repo/goapi/proto/$pack"
    echo "--go-grpc_opt=M$path=skywalking.apache.org/repo/goapi/proto/$pack"
  elif [[ $TYPE == "files" ]]; then
    echo "$PROTODIR/$name/$path"
  fi
}

proto envoy envoy/annotations/deprecation.proto
proto envoy envoy/api/v2/core/address.proto
proto envoy envoy/api/v2/core/backoff.proto
proto envoy envoy/api/v2/core/base.proto
proto envoy envoy/api/v2/core/config_source.proto
proto envoy envoy/api/v2/core/event_service_config.proto
proto envoy envoy/api/v2/core/grpc_method_list.proto
proto envoy envoy/api/v2/core/grpc_service.proto
proto envoy envoy/api/v2/core/http_uri.proto
proto envoy envoy/api/v2/core/socket_option.proto
proto envoy envoy/config/core/v3/address.proto
proto envoy envoy/config/core/v3/backoff.proto
proto envoy envoy/config/core/v3/base.proto
proto envoy envoy/config/core/v3/http_uri.proto
proto envoy envoy/config/core/v3/socket_option.proto
proto envoy envoy/data/accesslog/v2/accesslog.proto
proto envoy envoy/data/accesslog/v3/accesslog.proto
proto envoy envoy/service/accesslog/v2/als.proto
proto envoy envoy/service/accesslog/v3/als.proto
proto envoy envoy/service/metrics/v2/metrics_service.proto
proto envoy envoy/service/metrics/v3/metrics_service.proto
proto envoy envoy/type/percent.proto
proto envoy envoy/type/semantic_version.proto
proto envoy envoy/type/v3/percent.proto
proto envoy envoy/type/v3/semantic_version.proto
proto prometheus-model io/prometheus/client/metrics.proto
proto protoc-gen-validate validate/validate.proto
proto xds udpa/annotations/migrate.proto
proto xds udpa/annotations/status.proto
proto xds udpa/annotations/versioning.proto
proto xds xds/core/v3/context_params.proto
proto xds udpa/annotations/sensitive.proto

proto opentelementry opentelemetry/proto/collector/metrics/v1/metrics_service.proto
proto opentelementry opentelemetry/proto/metrics/v1/metrics.proto
proto opentelementry opentelemetry/proto/common/v1/common.proto
proto opentelementry opentelemetry/proto/resource/v1/resource.proto