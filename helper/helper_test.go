package helper

import "testing"

func TestNormalize(t *testing.T) {
	lowerCaseString := Normalize("zoë")

	if lowerCaseString != "zoe" {
		t.Errorf("Normalization error on lowerCase, expected %s, got %s", "zoe", lowerCaseString)
	}

	upperCaseString := Normalize("ZOË")

	if upperCaseString != "zoe" {
		t.Errorf("Normalization error on upperCase, expected %s got %s", "zoe", upperCaseString)
	}

	normalString := Normalize("zoe")

	if normalString != "zoe" {
		t.Errorf("Normalization error on normalString, expected %s, got %s", "zoe", normalString)
	}
}
