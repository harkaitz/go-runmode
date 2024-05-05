package runmode

import (
	"os"
)

// ReleaseMode is a flag to determine if the application is running in
// release mode. It is set to true if the environment variable RELEASE_MODE
// is set to any value.
//
// A program can set ReleaseMode to true if it is running in release mode.
var ReleaseMode bool = false

func init() {
	ReleaseMode = (os.Getenv("RELEASE_MODE") != "")
}

// IsReleaseMode returns true if the application is running in release mode.
func IsReleaseMode() bool {
	return ReleaseMode
}

// IsDebugMode returns true if the application is running in debug mode.
func IsDebugMode() bool {
	return !ReleaseMode
}
