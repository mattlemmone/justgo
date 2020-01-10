// +build !minimal

#pragma once

#ifndef GO_QTQUICKCONTROLS2_H
#define GO_QTQUICKCONTROLS2_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtQuickControls2_PackedString { char* data; long long len; void* ptr; };
struct QtQuickControls2_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif