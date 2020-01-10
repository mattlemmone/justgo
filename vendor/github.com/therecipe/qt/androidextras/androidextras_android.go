// +build android android_emulator

package androidextras

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "androidextras_android.h"
import "C"
import (
	"errors"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtAndroidExtras_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtAndroidExtras_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}
func QAndroidJniEnvironment_ExceptionCatch() error {
	var err error
	if QAndroidJniEnvironment_ExceptionCheck() {
		tmpExcPtr := QAndroidJniEnvironment_ExceptionOccurred()
		QAndroidJniEnvironment_ExceptionClear()
		tmpExc := NewQAndroidJniObject6(tmpExcPtr)
		err = errors.New(tmpExc.CallMethodString2("toString", "()Ljava/lang/String;"))
		tmpExc.DestroyQAndroidJniObject()
	}
	return err
}

func (e *QAndroidJniEnvironment) ExceptionCatch() error {
	return QAndroidJniEnvironment_ExceptionCatch()
}
func init() {
}
