// +build !minimal

#pragma once

#ifndef GO_QTDBUS_H
#define GO_QTDBUS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtDBus_PackedString { char* data; long long len; void* ptr; };
struct QtDBus_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif