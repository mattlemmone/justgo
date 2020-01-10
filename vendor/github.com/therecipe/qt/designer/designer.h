// +build !minimal

#pragma once

#ifndef GO_QTDESIGNER_H
#define GO_QTDESIGNER_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtDesigner_PackedString { char* data; long long len; void* ptr; };
struct QtDesigner_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif