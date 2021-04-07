/* SPDX-License-Identifier: BSD-2 */
package ioctl

import (
	"syscall"
	"testing"
	"unsafe"
)

func TestIoctl(t *testing.T) {
	const SIOCETHTOOL = 0x8946
	const ETHTOOL_GDRVINFO = 0x00000003
	type ifreq struct {
		ifr_name [16]byte
		ifr_data uintptr
	}
	type ethtoolDrvInfo struct {
		cmd          uint32
		driver       [32]byte
		version      [32]byte
		fw_version   [32]byte
		bus_info     [32]byte
		erom_version [32]byte
		reserved2    [12]byte
		n_priv_flags uint32
		n_stats      uint32
		testinfo_len uint32
		eedump_len   uint32
		regdump_len  uint32
	}
	drvinfo := ethtoolDrvInfo{
		cmd: ETHTOOL_GDRVINFO,
	}
	var name [16]byte
	copy(name[:], []byte("eth0"))
	ifr := ifreq{
		ifr_name: name,
		ifr_data: uintptr(unsafe.Pointer(&drvinfo)),
	}
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_IP)
	if err != nil {
		t.Errorf("unable to open socket file for ethtool")
	}
	ret := Ioctl(fd, SIOCETHTOOL, uintptr(unsafe.Pointer(&ifr)))
	if ret != nil {
		t.Errorf("Ioctl error %v", ret)
	}
	var stats_support string
	if drvinfo.n_stats != 0 {
		stats_support = "yes"
	} else {
		stats_support = "no"
	}
	t.Logf("%s info: \ndriver: %s\nversion: %s\nfirmware-version: %s\n"+
		"bus-info: %s\nsupports-statistics: %s\n",
		string(ifr.ifr_name[:]), string(drvinfo.driver[:]),
		string(drvinfo.version[:]), string(drvinfo.fw_version[:]),
		string(drvinfo.bus_info[:]), stats_support)
}
