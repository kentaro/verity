package verity

import (
	"os"
	"testing"
)

func TestCollector(t *testing.T) {
	os.Setenv("VERITY_TEST", "1")

	result, err := Collect()

	if err != nil {
		t.Errorf("got an error: %v\n", err)
	}

	if result["cpu"] == "foo" {
		t.Error("it should be able to collect CPU info\n")
	}

	if result["test"] == nil {
		t.Error("it should be able to collect environment variables\n")
	}

	if result["hostname"] == nil {
		t.Error("it should be able to collect hostname info\n")
	}

	if result["memory"] == nil {
		t.Error("it should be able to collect memory info\n")
	}
}
