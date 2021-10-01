#!/usr/bin/env bash

#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

PROTOC_VERSION=3.14.0

BASEDIR=$(dirname "$0")/..
TEMPDIR="$BASEDIR"/temp
BINDIR="$TEMPDIR"/bin
INCLUDE_DIR="$TEMPDIR"/include

mkdir -p "$TEMPDIR" || exit 1
mkdir -p "$BINDIR" || exit 1

if [[ -f "$BINDIR"/protoc ]]; then
  "$BINDIR"/protoc -I$INCLUDE_DIR/ "$@"
  exit $?
fi

if uname -s | grep MINGW64_NT || uname -s | grep CYGWIN_NT-6.1; then
  PROTOC_ZIP=protoc-"$PROTOC_VERSION"-win64.zip
elif uname -s | grep Darwin; then
  PROTOC_ZIP=protoc-"$PROTOC_VERSION"-osx-x86_64.zip
elif uname -s | grep Linux; then
  PROTOC_ZIP=protoc-"$PROTOC_VERSION"-linux-x86_64.zip
else
  echo "Sorry, we cannot install protoc for you, please visit https://github.com/protocolbuffers/protobuf and install protoc by yourself."
fi

curl -sL https://github.com/protocolbuffers/protobuf/releases/download/v"$PROTOC_VERSION"/$PROTOC_ZIP -o "$TEMPDIR"/$PROTOC_ZIP
unzip -o "$TEMPDIR"/$PROTOC_ZIP -d "$BINDIR"/.. bin/protoc > /dev/null 2>&1 || true
unzip -o "$TEMPDIR"/$PROTOC_ZIP -d "$BINDIR"/.. bin/protoc.exe > /dev/null 2>&1 || true
unzip -o "$TEMPDIR"/$PROTOC_ZIP -d "$BINDIR"/.. include/* > /dev/null 2>&1 || true

mv "$BINDIR"/protoc.exe "$BINDIR"/protoc > /dev/null 2>&1 || true

chmod +x "$BINDIR"/protoc

rm -f "$TEMPDIR"/$PROTOC_ZIP

"$BINDIR"/protoc -I$INCLUDE_DIR/ "$@"
