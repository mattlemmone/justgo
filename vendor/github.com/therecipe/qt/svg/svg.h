// +build !minimal

#pragma once

#ifndef GO_QTSVG_H
#define GO_QTSVG_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtSvg_PackedString { char* data; long long len; void* ptr; };
struct QtSvg_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif