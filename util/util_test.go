package util

import (
	"testing"
)

func TestPlatform(t *testing.T) {
	actual := Platform()
	expected := "linux"

	if actual != expected {
		t.Errorf("got %v\nexpected %v", actual, expected)
	}
}

func TestDistro(t *testing.T) {
	actual, err := Distro()
	expected := "rhel"

	if err != nil {
		t.Errorf("it should be able to retrieve distro name: %s", err)
	}

	if actual != expected {
		t.Errorf("got %v\nexpected %v", actual, expected)
	}
}
