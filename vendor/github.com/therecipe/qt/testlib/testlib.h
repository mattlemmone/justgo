// +build !minimal

#pragma once

#ifndef GO_QTTESTLIB_H
#define GO_QTTESTLIB_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtTestLib_PackedString { char* data; long long len; void* ptr; };
struct QtTestLib_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif