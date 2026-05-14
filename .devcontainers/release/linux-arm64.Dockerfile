# Copyright 2023 The Jule Project Contributors. All rights reserved.
# Use of this source code is governed by a BSD 3-Clause
# license that can be found in the LICENSE file.

FROM ubuntu:latest

RUN apt-get update -y
RUN apt-get install -y clang
RUN apt-get install -y ca-certificates
RUN apt-get install -y curl
RUN apt-get install -y p7zip-full

RUN mkdir /usr/local/workspace
WORKDIR /usr/local/workspace

RUN curl -L -o jule.zip https://github.com/julelang/jule/archive/refs/heads/master.zip
RUN 7z x jule.zip
RUN mv ./jule-master ./jule
WORKDIR /usr/local/workspace/jule

RUN mkdir ./bin
RUN curl -L -o ir.cpp https://raw.githubusercontent.com/julelang/julec-ir/main/src/linux-arm64.cpp
RUN curl -fsSL https://raw.githubusercontent.com/julelang/julec-ir/refs/heads/main/meta/linux-arm64.txt | bash

WORKDIR /usr/local/workspace
RUN 7z a -ttar -xr'!*.DS_Store' -xr'!__MACOSX' jule-linux-arm64.tar jule
RUN 7z a -txz jule-linux-arm64.tar.xz jule-linux-arm64.tar
RUN rm jule-linux-arm64.tar