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

func (self *Packages) Collect() (packages []string, err error) {
	packages, err = getInstalledPackages()
	return
}

func getInstalledPackages() (packages []string, err error) {
	distro, err := util.Distro()

	if err != nil {
		return
	}

	var command string
	var args []string

	switch distro {
	case "rhel":
		command = "rpm"
		args = []string{"-qa"}
	default:
		err = errors.New(fmt.Sprintf("unsupported distribution: %s", distro))
		return
	}

	output, err := exec.Command(command, args...).Output()

	if err != nil {
		return
	}

	packages = strings.Split(string(output), "\n")
	return
}
