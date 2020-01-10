// +build !minimal

#pragma once

#ifndef GO_QTLOCATION_H
#define GO_QTLOCATION_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtLocation_PackedString { char* data; long long len; void* ptr; };
struct QtLocation_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif