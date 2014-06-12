package system

import (
	"github.com/kentaro/verity/util"
	"os"
	"os/exec"
	"strings"
)

type System struct{}

const name = "system"

func (self *System) Name() string {
	return name
}

func (self *System) Collect() (result interface{}, err error) {
	result, err = getSystemInfo()
	return
}

func getSystemInfo() (info map[string]string, err error) {
	info = make(map[string]string)
	commands := map[string]string{
		"name":    "-s",
		"release": "-r",
		"version": "-v",
		"machine": "-m",
	}

	for key, arg := range commands {
		var out []byte

		cmd := exec.Command("uname", arg)
		cmd.Stderr = os.Stderr

		out, err = cmd.Output()
		if err != nil {
			return
		}

		info[key] = strings.TrimSpace(string(out))
	}

	var distro string

	distro, err = util.Distro()
	if err != nil {
		return
	}

	info["distribution"] = distro

	return
}
