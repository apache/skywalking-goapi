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

BASEDIR=$(dirname "$0")/..
TEMPDIR="$BASEDIR"/temp

. "$BASEDIR"/dependencies.sh

if [[ ! -d "$TEMPDIR" ]]; then
  mkdir -p "$TEMPDIR"
else
  rm -rf "${TEMPDIR:?}"/*
fi

curl -sLo "$TEMPDIR"/query-protocol.tgz https://github.com/apache/skywalking-query-protocol/archive/"${QUERY_PROTOCOL_SHA}".tar.gz

if [[ ! -d "$TEMPDIR"/query-protocol ]]; then
  mkdir "$TEMPDIR"/query-protocol
else
  rm -rf "$TEMPDIR"/query-protocol/*
fi

tar -zxf "$TEMPDIR"/query-protocol.tgz -C "$TEMPDIR"/query-protocol --strip 1

rm -rf "$TEMPDIR"/query-protocol.tgz

go get github.com/99designs/gqlgen

"$(go env GOPATH)"/bin/gqlgen -h > /dev/null 2>&1 || GO111MODULE=off go get github.com/99designs/gqlgen
go run "$BASEDIR"/scripts/tools/query_mutation.go

rm -rf "$TEMPDIR"/query-protocol

go mod tidy
