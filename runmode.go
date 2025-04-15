package runmode

import (
	"os"
	"strings"
)

// ReleaseMode is a flag to determine if the application is running in
// release mode. It is set to true if the environment variable RELEASE_MODE
// is set to any value.
//
// A program can set ReleaseMode to true if it is running in release mode.
var ReleaseMode bool = false

// Hostname holds the machine name (without domain).
var Hostname string = ""

func init() {
	var dot int
	var err error

	ReleaseMode = (os.Getenv("RELEASE_MODE") != "")

	Hostname, err = os.Hostname()
	if err != nil { panic(err) }
	dot = strings.Index(Hostname, ".")
	if dot != -1 {
		Hostname = Hostname[:dot]
	}
}

// IsReleaseMode returns true if the application is running in release mode.
func IsReleaseMode() bool {
	return ReleaseMode
}

// IsDebugMode returns true if the application is running in debug mode.
func IsDebugMode() bool {
	return !ReleaseMode
}

// DebugRelease returns first argument on debug second in release.
func DebugRelease(debug string, release string) string {
	switch IsDebugMode() {
	case true: return debug
	default:   return release
	}
}
