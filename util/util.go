package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func Platform() (result string) {
	result = runtime.GOOS
	return
}

func PackageFormat() (result string, err error) {
	platform := Platform()

	switch platform {
	case "linux":
		if _, err := os.Stat("/etc/redhat-release"); err == nil {
			result = "rpm"
		} else {
			err = errors.New("unsupported distribution")
		}
	case "darwin":
		result = "brew"
	default:
		err = errors.New(fmt.Sprintf("unsupported platform: %s", platform))
	}

	return
}

func Distro() (result string, err error) {
	format, err := PackageFormat()

	if err != nil {
		return
	}

	switch format {
	case "rpm":
		var data []byte

		data, err = ioutil.ReadFile("/etc/redhat-release")
		if err != nil {
			return
		}

		result = strings.TrimRight(string(data), "\n")
	case "brew":
		result = "osx"
	default:
		err = errors.New(fmt.Sprint("unsupported package format: %s", format))
	}

	return
}
