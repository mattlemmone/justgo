// +build !minimal

#pragma once

#ifndef GO_QTQUICK_H
#define GO_QTQUICK_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtQuick_PackedString { char* data; long long len; void* ptr; };
struct QtQuick_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif