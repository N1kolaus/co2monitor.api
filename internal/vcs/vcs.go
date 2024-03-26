package vcs

import (
	"fmt"
	"runtime/debug"
)

func Version() string {
	var (
		modified bool
		revision string
		time     string
	)

	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range buildInfo.Settings {
			switch s.Key {
			case "vcs.modified":
				if s.Value == "true" {
					modified = true
				}
			case "vcs.revision":
				revision = s.Value
			case "vcs.time":
				time = s.Value
			}

		}
	}

	if modified {
		return fmt.Sprintf("%s-%s-dirty", time, revision)
	}

	return fmt.Sprintf("%s-%s", time, revision)
}
