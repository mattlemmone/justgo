// +build !minimal

#pragma once

#ifndef GO_QTGAMEPAD_H
#define GO_QTGAMEPAD_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtGamepad_PackedString { char* data; long long len; void* ptr; };
struct QtGamepad_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif