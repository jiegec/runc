//go:build linux && (arm64 || amd64 || loong64 || mips || mipsle || mips64 || mips64le || ppc || ppc64 || ppc64le || riscv64 || s390x)
// +build linux
// +build arm64 amd64 loong64 mips mipsle mips64 mips64le ppc ppc64 ppc64le riscv64 s390x

package system

import (
	"golang.org/x/sys/unix"
)

// Setuid sets the uid of the calling thread to the specified uid.
func Setuid(uid int) (err error) {
	_, _, e1 := unix.RawSyscall(unix.SYS_SETUID, uintptr(uid), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

// Setgid sets the gid of the calling thread to the specified gid.
func Setgid(gid int) (err error) {
	_, _, e1 := unix.RawSyscall(unix.SYS_SETGID, uintptr(gid), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
