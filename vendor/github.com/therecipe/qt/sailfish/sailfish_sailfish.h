// +build sailfish sailfish_emulator

#pragma once

#ifndef GO_QTSAILFISH_H
#define GO_QTSAILFISH_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtSailfish_PackedString { char* data; long long len; void* ptr; };
struct QtSailfish_PackedList { void* data; long long len; };
int SailfishApp_SailfishApp_Main(int argc, char* argv);

#ifdef __cplusplus
}
#endif

#endif