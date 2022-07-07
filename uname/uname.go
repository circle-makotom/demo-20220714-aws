package uname

import (
	"fmt"
	"syscall"
)

func GetUnameString() (string, error) {
	var unameBuf syscall.Utsname

	err := syscall.Uname(&unameBuf)

	return fmt.Sprintf("%s %s %s %s", parseChars(unameBuf.Sysname), parseChars(unameBuf.Release), parseChars(unameBuf.Version), parseChars(unameBuf.Machine)), err
}

func parseChars(chars [65]int8) string {
	ret := ""

	for _, chr := range chars {
		if chr == 0 {
			break
		}
		ret += fmt.Sprintf("%c", chr)
	}

	return ret
}
