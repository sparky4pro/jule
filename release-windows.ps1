# Copyright 2026 The Jule Project Contributors. All rights reserved.
# Use of this source code is governed by a BSD 3-Clause
# license that can be found in the LICENSE file.

# Create a workspace to build a release.
New-Item -ItemType Directory -Path ".\release-workspace" -Force
Set-Location -Path ".\release-workspace"

# Get source code from the master branch and setup the environment.
curl.exe -L -o jule.zip https://github.com/julelang/jule/archive/refs/heads/master.zip
7z.exe x jule.zip
Remove-Item ".\jule.zip" -Recurse -Force
Rename-Item -Path ".\jule-master" -NewName "jule"
Set-Location -Path ".\jule"
New-Item -ItemType Directory -Path ".\bin" -Force

# Get ARM64 IRs from the GitHub and compile it.
curl.exe -L -o ir.cpp https://raw.githubusercontent.com/julelang/julec-ir/main/src/windows-arm64.cpp
clang++ -static -Wno-everything --std=c++20 -fwrapv -ffloat-store -fno-fast-math -fno-rounding-math -ffp-contract=fast -O3 -flto=thin -fuse-ld=lld -DNDEBUG -fomit-frame-pointer -fno-strict-aliasing -o .\bin\julec.exe ir.cpp -lws2_32 -lshell32 -liphlpapi -lsynchronization

# ZIP it.
Set-Location -Path ".."
7z.exe a -tzip -xr'!*.DS_Store' -xr'!__MACOSX' jule-windows-arm64.zip jule

# Clean environment for AMD64.
Set-Location -Path ".\jule"
Remove-Item ".\bin\*" -Recurse -Force
Remove-Item ".\ir.cpp"

# Get AMD64 IRs from the GitHub and compile it.
curl.exe -L -o ir.cpp https://raw.githubusercontent.com/julelang/julec-ir/main/src/windows-amd64.cpp
x86_64-w64-mingw32-clang++ -static -Wno-everything --std=c++20 -fwrapv -ffloat-store -fno-fast-math -fno-rounding-math -ffp-contract=fast -O3 -flto=thin -fuse-ld=lld -DNDEBUG -fomit-frame-pointer -fno-strict-aliasing -o .\bin\julec.exe ir.cpp -lws2_32 -lshell32 -liphlpapi -lsynchronization

# ZIP it.
Set-Location -Path ".."
7z.exe a -tzip -xr'!*.DS_Store' -xr'!__MACOSX' jule-windows-amd64.zip jule

# Clear the workspace.
Remove-Item ".\jule" -Recurse -Force