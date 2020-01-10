// +build !minimal

#pragma once

#ifndef GO_QTCHARTS_H
#define GO_QTCHARTS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtCharts_PackedString { char* data; long long len; void* ptr; };
struct QtCharts_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif