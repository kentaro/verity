package memory

import (
	"testing"
)

func TestMemory(t *testing.T) {
	collector := &Memory{}
	result, err := collector.Collect()

	if err != nil {
		t.Error("it should be able to collect hostname\n")
	}

	{
		actual := collector.Name()
		expected := "memory"

		if actual != expected {
			t.Errorf("got %v\nexpected %v\n", actual, expected)
		}
	}

	if result.(map[string]string)["total"] == "" {
		t.Error("it should be able to collect total memory size\n")
	}
}
