package main

import (
	"github.com/blang/semver"
	"testing"
)

var originalVersion = "1.1.1"

func TestSetMinorVersion(t *testing.T) {
	expected := "1.10.1"
	version, _ := semver.Make(originalVersion)
	newVersion := updateVersionNumber(version, "set", "minor", 10)

	if newVersion.String() != expected {
		t.Errorf("Got unexpected version, expected %s, got %s", expected, newVersion.String())
	}
}

func TestIncrementMinorVersion(t *testing.T) {
	expected := "1.2.1"
	version, _ := semver.Make(originalVersion)
	newVersion := updateVersionNumber(version, "increment", "minor", 1)

	if newVersion.String() != expected {
		t.Errorf("Got unexpected version, expected %s, got %s", expected, newVersion.String())
	}
}

func TestDecrementMinorVersion(t *testing.T) {
	expected := "1.0.1"
	version, _ := semver.Make(originalVersion)
	newVersion := updateVersionNumber(version, "decrement", "minor", 1)

	if newVersion.String() != expected {
		t.Errorf("Got unexpected version, expected %s, got %s", expected, newVersion.String())
	}
}

func TestSetMajorVersion(t *testing.T) {
	expected := "10.1.1"
	version, _ := semver.Make(originalVersion)
	newVersion := updateVersionNumber(version, "set", "major", 10)

	if newVersion.String() != expected {
		t.Errorf("Got unexpected version, expected %s, got %s", expected, newVersion.String())
	}
}

func TestIncrementMajorVersion(t *testing.T) {
	expected := "2.1.1"
	version, _ := semver.Make(originalVersion)
	newVersion := updateVersionNumber(version, "increment", "major", 1)

	if newVersion.String() != expected {
		t.Errorf("Got unexpected version, expected %s, got %s", expected, newVersion.String())
	}
}

func TestDecrementMajorVersion(t *testing.T) {
	expected := "0.1.1"
	version, _ := semver.Make(originalVersion)
	newVersion := updateVersionNumber(version, "decrement", "major", 1)

	if newVersion.String() != expected {
		t.Errorf("Got unexpected version, expected %s, got %s", expected, newVersion.String())
	}
}

func TestSetPatchVersion(t *testing.T) {
	expected := "1.1.10"
	version, _ := semver.Make(originalVersion)
	newVersion := updateVersionNumber(version, "set", "patch", 10)

	if newVersion.String() != expected {
		t.Errorf("Got unexpected version, expected %s, got %s", expected, newVersion.String())
	}
}

func TestIncrementPatchVersion(t *testing.T) {
	expected := "1.1.2"
	version, _ := semver.Make(originalVersion)
	newVersion := updateVersionNumber(version, "increment", "patch", 1)

	if newVersion.String() != expected {
		t.Errorf("Got unexpected version, expected %s, got %s", expected, newVersion.String())
	}
}

func TestDecrementPatchVersion(t *testing.T) {
	expected := "1.1.0"
	version, _ := semver.Make(originalVersion)
	newVersion := updateVersionNumber(version, "decrement", "patch", 1)

	if newVersion.String() != expected {
		t.Errorf("Got unexpected version, expected %s, got %s", expected, newVersion.String())
	}
}
