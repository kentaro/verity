package cpu

type Cpu struct{}

const name = "cpu"

func (self *Cpu) Name() string {
	return name
}
