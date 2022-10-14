package gouilded

import (
	"runtime"
	"runtime/debug"
	"strings"
)

const (
	// Name is the library name
	Name = "gouilded"

	// GitHub is a link to the libraries GitHub repository
	GitHub = "https://github.com/yyewolf/" + Name
)

var (
	// Version is the currently used version of Gouilded
	Version = getVersion()

	// OS is the currently used OS
	OS = getOS()
)

func getVersion() string {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, dep := range bi.Deps {
			if strings.Contains(GitHub, dep.Path) {
				return dep.Version
			}
		}
	}
	return "unknown"
}

func getOS() string {
	os := runtime.GOOS
	if strings.HasPrefix(os, "windows") {
		return "windows"
	}
	if strings.HasPrefix(os, "darwin") {
		return "darwin"
	}
	return "linux"
}
