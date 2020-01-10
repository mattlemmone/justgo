// +build !minimal

#pragma once

#ifndef GO_QTSERIALPORT_H
#define GO_QTSERIALPORT_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtSerialPort_PackedString { char* data; long long len; void* ptr; };
struct QtSerialPort_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif