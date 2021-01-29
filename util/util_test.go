package util

import "testing"

func TestNormalize(t *testing.T) {
	lowerCaseString := Normalize("zoë")

	if lowerCaseString != "zoe" {
		t.Errorf("Normalization error on lowerCase - expected %s, got %s", "zoe", lowerCaseString)
	}

	upperCaseString := Normalize("ZOË")

	if upperCaseString != "zoe" {
		t.Errorf("Normalization error on upperCase - expected %s got %s", "zoe", upperCaseString)
	}

	normalString := Normalize("zoe")

	if normalString != "zoe" {
		t.Errorf("Normalization error on normalString - expected %s, got %s", "zoe", normalString)
	}
}

func TestNormalizeMapKeys(t *testing.T) {
	denormalizedMap := make(map[string]string)

	denormalizedMap["UPPERCASE"] = "_"
	denormalizedMap["zoë"] = "_"

	NormalizeMapKeys(denormalizedMap)

	if _, ok := denormalizedMap["zoe"]; !ok {
		t.Error("MapKeyNormalization error on zoë")
	}

	if _, ok := denormalizedMap["uppercase"]; !ok {
		t.Error("MapKeyNormalization error on UPPERCASE")
	}

	if _, ok := denormalizedMap["UPPERCASE"]; ok {
		t.Error("MapKeyNormalization error - should have deleted key UPPERCASE")
	}

	if _, ok := denormalizedMap["zoë"]; ok {
		t.Error("MapKeyNormalization error - should have deleted key zoë")
	}
}
