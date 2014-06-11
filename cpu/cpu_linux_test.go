package cpu

import (
	"testing"
)

func TestCpu(t *testing.T) {
	collector := &Cpu{}
	result, err := collector.Collect()

	if err != nil {
		t.Error("it should be able to collect CPU info\n")
	}

	{
		actual := collector.Name()
		expected := "cpu"

		if actual != expected {
			t.Errorf("got %v\nexpected %v\n", actual, expected)
		}
	}

	if result.(map[string]string)["model_name"] == "" {
		t.Error("it should be able to collect CPU model name\n")
	}

	if result.(map[string]string)["total"] == "0" {
		t.Error("it should be able to collect the number of CPU(s)\n")
	}
}
