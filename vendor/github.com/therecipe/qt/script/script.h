// +build !minimal

#pragma once

#ifndef GO_QTSCRIPT_H
#define GO_QTSCRIPT_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtScript_PackedString { char* data; long long len; void* ptr; };
struct QtScript_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif