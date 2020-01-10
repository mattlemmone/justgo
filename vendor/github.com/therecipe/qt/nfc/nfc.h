// +build !minimal

#pragma once

#ifndef GO_QTNFC_H
#define GO_QTNFC_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtNfc_PackedString { char* data; long long len; void* ptr; };
struct QtNfc_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif