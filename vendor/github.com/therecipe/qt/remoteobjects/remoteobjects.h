// +build !minimal

#pragma once

#ifndef GO_QTREMOTEOBJECTS_H
#define GO_QTREMOTEOBJECTS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtRemoteObjects_PackedString { char* data; long long len; void* ptr; };
struct QtRemoteObjects_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif