// +build !minimal

#pragma once

#ifndef GO_QTSCRIPTTOOLS_H
#define GO_QTSCRIPTTOOLS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtScriptTools_PackedString { char* data; long long len; void* ptr; };
struct QtScriptTools_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif