// +build !minimal

#pragma once

#ifndef GO_QTWEBENGINE_H
#define GO_QTWEBENGINE_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtWebEngine_PackedString { char* data; long long len; void* ptr; };
struct QtWebEngine_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif