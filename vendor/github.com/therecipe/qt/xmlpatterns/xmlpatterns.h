// +build !minimal

#pragma once

#ifndef GO_QTXMLPATTERNS_H
#define GO_QTXMLPATTERNS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtXmlPatterns_PackedString { char* data; long long len; void* ptr; };
struct QtXmlPatterns_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif