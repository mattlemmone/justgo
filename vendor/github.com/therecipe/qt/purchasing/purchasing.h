// +build !minimal

#pragma once

#ifndef GO_QTPURCHASING_H
#define GO_QTPURCHASING_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtPurchasing_PackedString { char* data; long long len; void* ptr; };
struct QtPurchasing_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif