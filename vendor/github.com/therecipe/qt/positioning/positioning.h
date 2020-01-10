// +build !minimal

#pragma once

#ifndef GO_QTPOSITIONING_H
#define GO_QTPOSITIONING_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtPositioning_PackedString { char* data; long long len; void* ptr; };
struct QtPositioning_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif