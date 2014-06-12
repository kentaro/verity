package packages

import (
	"testing"
)

func TestPackages(t *testing.T) {
	collector := &Packages{}
	_, err := collector.Collect()

	if err != nil {
		t.Error("it should be able to collect installed packages\n")
	}

	{
		actual := collector.Name()
		expected := "packages"

		if actual != expected {
			t.Errorf("got %v\nexpected %v\n", actual, expected)
		}
	}
}

