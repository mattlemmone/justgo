// +build !minimal

#pragma once

#ifndef GO_QTPRINTSUPPORT_H
#define GO_QTPRINTSUPPORT_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtPrintSupport_PackedString { char* data; long long len; void* ptr; };
struct QtPrintSupport_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif