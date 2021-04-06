/* SPDX-License-Identifier: BSD-2 */
package ioctl

import (
	"syscall"
)

/* definitions from arch/ioctl.h */

const (
	/*bits len*/
	ioc_nrbits   = 8
	ioc_typebits = 8
	ioc_sizebits = 14
	ioc_dirbits  = 2
	/*mask bits*/
	ioc_nrmask   = (1 << ioc_nrbits) - 1
	ioc_typemask = (1 << ioc_typebits) - 1
	ioc_sizemask = (1 << ioc_sizebits) - 1
	ioc_dirmask  = (1 << ioc_dirbits) - 1
	/*shift bits*/
	ioc_nrshift   = 0
	ioc_typeshift = ioc_nrshift + ioc_nrbits
	ioc_sizeshift = ioc_typeshift + ioc_typebits
	ioc_dirshift  = ioc_sizeshift + ioc_sizebits

	/*io action*/
	IO_NONE  = 0
	IO_READ  = 1
	IO_WRITE = 2
)

func ioc(dir, tp, nr, size uint) uint {
	return dir<<ioc_dirshift | tp<<ioc_typeshift | nr<<ioc_nrshift | size<<ioc_sizeshift
}

// Io
func Io(tp, nr uint) uint {
	return ioc(IO_NONE, tp, nr, 0)
}

//Ior
func Ior(tp, nr uint) uint {
	return ioc(IO_READ, tp, nr, 0)
}

// Iow
func Iow(tp, nr, size uint) uint {
	return ioc(IO_WRITE, tp, nr, size)
}

// Iowr
func Iowr(tp, nr, size uint) uint {
	return ioc(IO_READ|IO_WRITE, tp, nr, size)
}

// Ior_bad
func Ior_bad(tp, nr, size uint) uint {
	return ioc(IO_WRITE, tp, nr, size)
}

//Iow_bad
func Iow_bad(tp, nr, size uint) uint {
	return ioc(IO_READ, tp, nr, size)
}

func Iowr_bad(tp, nr, size uint) uint {
	return ioc(IO_READ|IO_WRITE, tp, nr, size)
}

/* a list of request and related constants */

/* define some types used by request */

/* On unix system, defines some symbols for tty */

// Ioctl return ioctl result
func Ioctl(fd int, request uint64, arg ...uintptr) int {
	//in most case syscall 4 would fit well, just judge the input len
	//arg could be a integer value or a pointer to data(going to or coming from driver)
	if len(arg) == 0 {
		syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(request), 0)
	} else if len(arg) == 1 {
		syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(request), arg[0])
	} else if len(arg) == 4 {
		syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(request), arg[0], arg[1], arg[2], arg[3])
	}

	return 0
}

const (
	O_ACCMODE   = syscall.O_ACCMODE
	O_APPEND    = syscall.O_APPEND
	O_ASYNC     = syscall.O_ASYNC
	O_CLOEXEC   = syscall.O_CLOEXEC
	O_CREAT     = syscall.O_CREAT
	O_DIRECT    = syscall.O_DIRECT
	O_DIRECTORY = syscall.O_DIRECTORY
	O_DSYNC     = syscall.O_DSYNC
	O_EXCL      = syscall.O_EXCL
	O_FSYNC     = syscall.O_FSYNC
	O_LARGEFILE = syscall.O_LARGEFILE
	O_NDELAY    = syscall.O_NDELAY
	O_NOATIME   = syscall.O_NOATIME
	O_NOCTTY    = syscall.O_NOCTTY
	O_NOFOLLOW  = syscall.O_NOFOLLOW
	O_NONBLOCK  = syscall.O_NONBLOCK
	O_RDONLY    = syscall.O_RDONLY
	O_RDWR      = syscall.O_RDWR
	O_RSYNC     = syscall.O_RSYNC
	O_SYNC      = syscall.O_SYNC
	O_TRUNC     = syscall.O_TRUNC
	O_WRONLY    = syscall.O_WRONLY
)

func Open(path string, mode int, perm uint32) (int, error) {
	return syscall.Open(path, mode|syscall.O_NONBLOCK, perm)
}

func Close(fd int) error {
	return syscall.Close(fd)
}
