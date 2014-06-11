package system

import (
	"github.com/kentaro/verity/util"
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
	commands := map[string][]string{
		"name":    []string{"uname", "-s"},
		"release": []string{"uname", "-r"},
		"version": []string{"uname", "-v"},
		"machine": []string{"uname", "-m"},
		"os":      []string{"uname", "-o"},
	}

	for key, command := range commands {
		var out []byte
		out, err = exec.Command(command[0], command[1]).Output()
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
