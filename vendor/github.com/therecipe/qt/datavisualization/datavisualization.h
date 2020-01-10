// +build !minimal

#pragma once

#ifndef GO_QTDATAVISUALIZATION_H
#define GO_QTDATAVISUALIZATION_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtDataVisualization_PackedString { char* data; long long len; void* ptr; };
struct QtDataVisualization_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif