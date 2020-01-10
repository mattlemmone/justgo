// +build !minimal

#pragma once

#ifndef GO_QTBLUETOOTH_H
#define GO_QTBLUETOOTH_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtBluetooth_PackedString { char* data; long long len; void* ptr; };
struct QtBluetooth_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif