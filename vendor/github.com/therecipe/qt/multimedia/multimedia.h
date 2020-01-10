// +build !minimal

#pragma once

#ifndef GO_QTMULTIMEDIA_H
#define GO_QTMULTIMEDIA_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtMultimedia_PackedString { char* data; long long len; void* ptr; };
struct QtMultimedia_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif