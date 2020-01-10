// +build !minimal

#pragma once

#ifndef GO_QTSCXML_H
#define GO_QTSCXML_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtScxml_PackedString { char* data; long long len; void* ptr; };
struct QtScxml_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif