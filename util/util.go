package util

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

func Platform() (result string) {
	result = runtime.GOOS
	return
}

func Distro() (result string, err error) {
	platform := Platform()

	switch platform {
	case "linux":
		if _, err := os.Stat("/etc/redhat-release"); err == nil {
			result = "rhel"
		} else {
			err = errors.New("unsupported distribution")
		}
	default:
		err = errors.New(fmt.Sprint("unsupported platform: %s", platform))
	}

	return
}
