# Copyright 2026 The Jule Project Contributors. All rights reserved.
# Use of this source code is governed by a BSD 3-Clause
# license that can be found in the LICENSE file.

# Create a workspace to build a release.
mkdir ./jule-workspace
cd ./jule-workspace

# Get source code from the master branch and setup the environment.
curl -L -o jule.zip https://github.com/julelang/jule/archive/refs/heads/master.zip
7zz x ./jule.zip
rm ./jule.zip
mv ./jule-master ./jule
cd ./jule
mkdir ./bin

# Get ARM64 IRs from the GitHub and compile it.
curl -L -o ir.cpp https://raw.githubusercontent.com/julelang/julec-ir/main/src/darwin-arm64.cpp
clang++ -Wno-everything --std=c++20 -fwrapv -ffloat-store -fno-fast-math -fexcess-precision=standard -fno-rounding-math -ffp-contract=fast -O3 -flto=thin -DNDEBUG -fomit-frame-pointer -fno-strict-aliasing -o ./bin/julec ir.cpp

# Create tar.xz for release.
cd ..
7zz a -ttar -xr'!*.DS_Store' -xr'!__MACOSX' jule-darwin-arm64.tar jule
7zz a -txz jule-darwin-arm64.tar.xz jule-darwin-arm64.tar
rm jule-darwin-arm64.tar

# Clean environment for AMD64.
cd ./jule
rm ./bin/julec
rm ./ir.cpp

# Get AMD64 IRs from the GitHub and compile it.
curl -L -o ir.cpp https://raw.githubusercontent.com/julelang/julec-ir/main/src/darwin-amd64.cpp
clang++ -arch x86_64 -Wno-everything --std=c++20 -fwrapv -ffloat-store -fno-fast-math -fexcess-precision=standard -fno-rounding-math -ffp-contract=fast -O3 -flto=thin -DNDEBUG -fomit-frame-pointer -fno-strict-aliasing -o ./bin/julec ir.cpp

# Create tar.xz for release.
cd ..
7zz a -ttar -xr'!*.DS_Store' -xr'!__MACOSX' jule-darwin-amd64.tar jule
7zz a -txz jule-darwin-amd64.tar.xz jule-darwin-amd64.tar
rm jule-darwin-amd64.tar

# Clear the workspace.
rm -r ./jule