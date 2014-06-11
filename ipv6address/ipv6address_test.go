package ipv6address

import (
	"testing"
)

func TestIpv6Address(t *testing.T) {
	collector := &Ipv6Address{}
	_, err := collector.Collect()

	if err != nil {
		t.Error("it should be able to collect IP V6 address\n")
	}

	{
		actual := collector.Name()
		expected := "ipv6address"

		if actual != expected {
			t.Errorf("got %v\nexpected %v\n", actual, expected)
		}
	}
}
