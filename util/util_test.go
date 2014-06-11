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

func TestPackageFormat(t *testing.T) {
	actual, err := PackageFormat()
	expected := "rpm"

	if err != nil {
		t.Errorf("it should be able to retrieve package format: %s", err)
	}

	if actual != expected {
		t.Errorf("got %v\nexpected %v", actual, expected)
	}
}

func TestDistro(t *testing.T) {
	actual, err := Distro()
	expected := "CentOS release 6.5 (Final)"

	if err != nil {
		t.Errorf("it should be able to retrieve package format: %s", err)
	}

	if actual != expected {
		t.Errorf("got %v\nexpected %v", actual, expected)
	}
}
