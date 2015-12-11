// Copyright (c) 2012 VMware, Inc.

// +build darwin freebsd linux netbsd openbsd

package sigar

/*
#include <sys/utsname.h>
#include <unistd.h>
*/
import "C"

import (
	"io/ioutil"
	"os"
	"regexp"
	"syscall"
)

func (self *FileSystemUsage) Get(path string) error {
	stat := syscall.Statfs_t{}
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return err
	}

	bsize := stat.Bsize / 512

	self.Total = (uint64(stat.Blocks) * uint64(bsize)) >> 1
	self.Free = (uint64(stat.Bfree) * uint64(bsize)) >> 1
	self.Avail = (uint64(stat.Bavail) * uint64(bsize)) >> 1
	self.Used = self.Total - self.Free
	self.Files = stat.Files
	self.FreeFiles = stat.Ffree

	return nil
}


func (self *SystemInfo) getFromUname() {
	var unameBuf C.struct_utsname
	C.uname(&unameBuf)

	self.Version = C.GoString(&unameBuf.release[0])
	self.VendorName = C.GoString(&unameBuf.sysname[0])
	self.Name = C.GoString(&unameBuf.sysname[0])
	self.Machine = C.GoString(&unameBuf.machine[0])
	self.Arch = C.GoString(&unameBuf.machine[0])
	self.PatchLevel = "unknown"
}

func (self *NetworkInfo) GetForUnix() error {
	resolvFilePath := "/etc/resolv.conf"

	_, err := os.Stat(resolvFilePath)

	if os.IsNotExist(err) == false {
		resolvFile, _ := ioutil.ReadFile(resolvFilePath)

		regex, _ := regexp.Compile(`nameserver\ (.+)`)

		matches := regex.FindAllString(string(resolvFile), -1)

		for _, v := range matches {
			if self.PrimaryDns == "" {
				self.PrimaryDns = v
			} else if self.SecondaryDns == "" {
				self.SecondaryDns = v
			}

		}

	} else {
		panic(err)

		return nil
	}

	self.HostName, _ = os.Hostname()

	return nil
}
