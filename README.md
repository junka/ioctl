# ioctl

the syscall ```ioctl()``` maniplates the underlay deviceparameters of special files.

see the manpage of ioctl to read more.

this project aims to implements ioctl in golang like
```
func Ioctl(fd int, request uint64, arg ...uintptr) error
```
