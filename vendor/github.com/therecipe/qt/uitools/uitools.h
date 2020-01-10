// +build !minimal

#pragma once

#ifndef GO_QTUITOOLS_H
#define GO_QTUITOOLS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtUiTools_PackedString { char* data; long long len; void* ptr; };
struct QtUiTools_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif