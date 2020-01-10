// +build !minimal

#pragma once

#ifndef GO_QTWEBCHANNEL_H
#define GO_QTWEBCHANNEL_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtWebChannel_PackedString { char* data; long long len; void* ptr; };
struct QtWebChannel_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif