// +build android android_emulator

#pragma once

#ifndef GO_QTANDROIDEXTRAS_H
#define GO_QTANDROIDEXTRAS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtAndroidExtras_PackedString { char* data; long long len; void* ptr; };
struct QtAndroidExtras_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif