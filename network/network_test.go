package network

import (
	"testing"
)

func TestNetwork(t *testing.T) {
	collector := &Network{}
	_, err := collector.Collect()

	if err != nil {
		t.Error("it should be able to collect network interfaces\n")
	}

	{
		actual := collector.Name()
		expected := "network"

		if actual != expected {
			t.Errorf("got %v\nexpected %v\n", actual, expected)
		}
	}
}
