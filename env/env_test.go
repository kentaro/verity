package env

import (
	"os"
	"testing"
)

func TestHostname(t *testing.T) {
	os.Setenv("VERITY_TEST1", "1")
	os.Setenv("VERITY_TEST2", "2")

	collector := &Env{}
	result, err := collector.Collect()

	if err != nil {
		t.Error("it should be able to collect environment variables\n")
	}

	{
		actual := collector.Name()
		expected := "env"

		if actual != expected {
			t.Errorf("got %v\nexpected %v\n", actual, expected)
		}
	}

	{
		actual := result.(map[string]string)["test1"]
		expected := "1"

		if actual != expected {
			t.Errorf("got %v\nexpected %v\n", actual, expected)
		}
	}

	{
		actual := result.(map[string]string)["test2"]
		expected := "2"

		if actual != expected {
			t.Errorf("got %v\nexpected %v\n", actual, expected)
		}
	}
}
