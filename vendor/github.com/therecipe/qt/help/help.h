// +build !minimal

#pragma once

#ifndef GO_QTHELP_H
#define GO_QTHELP_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtHelp_PackedString { char* data; long long len; void* ptr; };
struct QtHelp_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif