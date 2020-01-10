// +build !minimal

#pragma once

#ifndef GO_QTSERIALBUS_H
#define GO_QTSERIALBUS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtSerialBus_PackedString { char* data; long long len; void* ptr; };
struct QtSerialBus_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif