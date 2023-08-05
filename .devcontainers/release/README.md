# Release Containers

This region is reserved for release periods and serves JuleC's build for different platforms.
Please use `build.sh` to run the whole process automatically.
The `build.sh` designed for root directory of project.
Please execute `build.sh` when you are in root directory of project.

## Supported Platforms

Listed below are the unique OS-ARCH operations used in the build process. \
The C++ IR file names that these matches should have during compilation are given opposite them.

- `linux-amd64`: `linux-amd64.cpp`
- `linux-arm64`: `linux-arm64.cpp`

## Preparing to Build

It is recommended that you first be in the [`root`](https://github.com/julelang/jule) directory of project.
You need to create a C++ IR for the respective platforms using JuleC's cross-transpilation with your latest compiler and name it as needed.
The nomenclatures are given in the system mappings above.

However, for a correct compilation you need to get rid of absolute include paths.
In absolute paths, you need to delete up to the root directory of the compiler.
Then correctly, you need to change directories to be portable.
If you transpiled in the recommended [`root`](https://github.com/julelang/jule) directory, your IR files are expected to be in the `dist` directory.
In this case, remove root directory of the compiler and append `../` prefix.

For example:

Absolute paths:

```cpp
#include "/foo/bar/jule/api/jule.hpp"

#include <dirent.h>
#include <fcntl.h>
#include <cstdio>
#include <sys/stat.h>
#include "/foo/bar/jule/std/sys/syscall_unix.hpp"
#include "/foo/bar/jule/std/os/proc.hpp"
#include "/foo/bar/jule/std/vector/alloc.hpp"
#include "/foo/bar/jule/std/jule/parser/parser.hpp"
#include "/foo/bar/jule/src/julec/obj/cxx/cxx.hpp"
#include "/foo/bar/jule/src/julec/julec.hpp"
```

Portable paths:

```cpp
#include "../api/jule.hpp"

#include <dirent.h>
#include <fcntl.h>
#include <cstdio>
#include <sys/stat.h>
#include "../std/sys/syscall_unix.hpp"
#include "../std/os/proc.hpp"
#include "../std/vector/alloc.hpp"
#include "../std/jule/parser/parser.hpp"
#include "../src/julec/obj/cxx/cxx.hpp"
#include "../src/julec/julec.hpp"
```