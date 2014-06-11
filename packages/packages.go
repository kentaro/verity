package packages

import (
	"errors"
	"fmt"
	"github.com/kentaro/verity/util"
	"os/exec"
	"strings"
)

type Packages struct{}

const name = "packages"

func (self *Packages) Name() string {
	return name
}

func (self *Packages) Collect() (result interface{}, err error) {
	result, err = getInstalledPackages()
	return
}

func getInstalledPackages() (packages []string, err error) {
	format, err := util.PackageFormat()

	if err != nil {
		return
	}

	var command string
	var args []string

	switch format {
	case "rpm":
		command = format
		args = []string{"-qa"}
	default:
		err = errors.New(fmt.Sprintf("unsupported distribution: %s", distro))
		return
	}

	output, err := exec.Command(command, args...).Output()

	if err != nil {
		return
	}

	packages = strings.Split(strings.TrimRight(string(output), "\n"), "\n")
	return
}
