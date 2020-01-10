// +build !minimal

#pragma once

#ifndef GO_QTXML_H
#define GO_QTXML_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtXml_PackedString { char* data; long long len; void* ptr; };
struct QtXml_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif