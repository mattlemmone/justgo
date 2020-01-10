// +build !minimal

#pragma once

#ifndef GO_QTVIRTUALKEYBOARD_H
#define GO_QTVIRTUALKEYBOARD_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtVirtualKeyboard_PackedString { char* data; long long len; void* ptr; };
struct QtVirtualKeyboard_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif