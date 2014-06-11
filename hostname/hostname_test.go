package hostname

import (
	"os"
	"testing"
)

func TestHostname(t *testing.T) {
	hostname, _ := os.Hostname()

	collector := &Hostname{}
	result, err := collector.Collect()

	if err != nil {
		t.Error("it should be able to collect hostname\n")
	}

	{
		actual := collector.Name()
		expected := "hostname"

		if actual != expected {
			t.Errorf("got %v\nexpected %v\n", actual, expected)
		}
	}

	{
		actual := result.(string)
		expected := hostname

		if actual != expected {
			t.Errorf("got %v\nexpected %v\n", actual, expected)
		}
	}
}
