//go:build unix

package renovatebugreproduction

import "golang.org/x/sys/unix"

const O_NONBLOCK = unix.O_NONBLOCK
