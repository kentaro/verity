package ipaddress

import (
	"testing"
)

func TestIpAddress(t *testing.T) {
	collector := &IpAddress{}
	_, err := collector.Collect()

	if err != nil {
		t.Error("it should be able to collect IP address\n")
	}

	{
		actual := collector.Name()
		expected := "ipaddress"

		if actual != expected {
			t.Errorf("got %v\nexpected %v\n", actual, expected)
		}
	}
}
