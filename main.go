package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/blang/semver"
	"os"
)

var (
	action = flag.String("action", "", "Action to do: set, increment, decrement")
	what   = flag.String("what", "patch", "Where to run the action: major, minor, patch")
	value  = flag.Uint64("value", 1, "The value to use during the action")
)

func updateVersionNumber(version semver.Version, action string, what string, value uint64) semver.Version {
	if action == "set" {
		if what == "major" {
			version.Major = value
		} else if what == "minor" {
			version.Minor = value
		} else if what == "patch" {
			version.Patch = value
		}
	} else if action == "increment" {
		if what == "major" {
			version.Major = version.Major + value
		} else if what == "minor" {
			version.Minor = version.Minor + value
		} else if what == "patch" {
			version.Patch = version.Patch + value
		}
	} else if action == "decrement" {
		if what == "major" {
			version.Major = version.Major - value
		} else if what == "minor" {
			version.Minor = version.Minor - value
		} else if what == "patch" {
			version.Patch = version.Patch - value
		}
	}
	return version
}

func main() {
	flag.Parse()
	source := bufio.NewScanner(os.Stdin)
	for source.Scan() {
		parsedVersion, _ := semver.Make(source.Text())
		fmt.Println(updateVersionNumber(parsedVersion, *action, *what, *value))

	}

}
