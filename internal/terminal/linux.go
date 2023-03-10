//go:build linux
// +build linux

package terminal

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TCGETS
const ioctlWriteTermios = unix.TCSETS
