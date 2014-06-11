package macaddress

import (
	"testing"
)

func TestMacAddress(t *testing.T) {
	collector := &MacAddress{}

	{
		actual := collector.Name()
		expected := "macaddress"

		if actual != expected {
			t.Errorf("got %v\nexpected %v\n", actual, expected)
		}
	}
}
