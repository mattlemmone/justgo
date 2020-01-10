// +build !android,!android_emulator

package androidextras

import (
	"strings"
)

func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}
func QAndroidJniEnvironment_ExceptionCatch() error {
	return nil
}

func (e *QAndroidJniEnvironment) ExceptionCatch() error { return nil }
func init() {
}
