package goversion

import "runtime"

func Version() string {
	return runtime.Version()
}
