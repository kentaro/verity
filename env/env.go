package env

import (
	"os"
	"strings"
)

type Env struct{}

const name = "env"

func (self *Env) Name() string {
	return name
}

func (self *Env) Collect() (result interface{}, err error) {
	result, err = self.getEnvironmentVariables()
	return
}

func (self *Env) getEnvironmentVariables() (envs map[string]string, err error) {
	envs = make(map[string]string)

	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if strings.HasPrefix(pair[0], "VERITY_") {
			key := strings.ToLower(strings.SplitAfterN(pair[0], "VERITY_", 2)[1])
			envs[key] = pair[1]
		}
	}

	return
}
