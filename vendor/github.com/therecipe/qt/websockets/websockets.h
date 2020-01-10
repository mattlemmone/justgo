// +build !minimal

#pragma once

#ifndef GO_QTWEBSOCKETS_H
#define GO_QTWEBSOCKETS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtWebSockets_PackedString { char* data; long long len; void* ptr; };
struct QtWebSockets_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif