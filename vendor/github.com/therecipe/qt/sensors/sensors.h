// +build !minimal

#pragma once

#ifndef GO_QTSENSORS_H
#define GO_QTSENSORS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtSensors_PackedString { char* data; long long len; void* ptr; };
struct QtSensors_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif