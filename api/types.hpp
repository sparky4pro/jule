#ifndef __JULE_TYPES_HPP
#define __JULE_TYPES_HPP

#include <stddef.h>

#include "platform.hpp"

namespace jule {

#ifdef ARCH_32BIT
    typedef unsigned long int Uint;
    typedef signed long int Int;
    typedef unsigned long int Uintptr;
#else
    typedef unsigned long long int Uint;
    typedef signed long long int Int;
    typedef unsigned long long int Uintptr;
#endif

    typedef signed char I8;
    typedef signed short int I16;
    typedef signed long int I32;
    typedef signed long long int I64;
    typedef unsigned char U8;
    typedef unsigned short int U16;
    typedef unsigned long int U32;
    typedef unsigned long long int U64;
    typedef float F32;
    typedef double F64;
    typedef bool Bool;

    constexpr decltype(nullptr) nil{ nullptr };

} // namespace jule

#endif // ifndef __JULE_TYPES_HPP