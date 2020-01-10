// +build !minimal

#pragma once

#ifndef GO_QTSQL_H
#define GO_QTSQL_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtSql_PackedString { char* data; long long len; void* ptr; };
struct QtSql_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif