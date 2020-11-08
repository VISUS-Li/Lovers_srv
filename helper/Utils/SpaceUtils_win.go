// +build windows

package Utils

import (
	"errors"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

//内存使用情况结构
type memoryStatusEx struct {
	dwLength                uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64 // in bytes
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

//磁盘使用情况结构
type DistStatus struct{
	DiskAll		int64
	DiskFree	int64
	DiskUsed	int64
	DiskAll_GB	float64
	DiskFree_GB	float64
	DiskUsed_GB	float64
}

var handle = syscall.MustLoadDLL("kernel32.dll")

/******
查询磁盘剩余空间
partition:系统为windows时，要查询的分区盘符，只传入盘符字母，如 C。
******/
func GetDiskSpaceStatus_win(partition string) (*DistStatus, error){
	sysType := runtime.GOOS
	if sysType == "windows"{
			//对partition参数进行过滤
		if len(partition) <= 0 || len(partition) > 1{
			partition = "C"
		}
		partition = strings.ToUpper(partition)
		partition = partition + ":"

		c := handle.MustFindProc("GetDiskFreeSpaceExW")
		lpFreeBytesAvailable := int64(0)
		lpTotalNumberOfBytes := int64(0)
		lpTotalNumberOfFreeBytes := int64(0)
		a1, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(partition))),
			uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
			uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
			uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
		if a1 == 0{
			logrus.Error("get disk space:" + err.Error())
		}
		diskStatus := new(DistStatus)
		diskStatus.DiskAll = lpTotalNumberOfBytes
		diskStatus.DiskFree = lpFreeBytesAvailable
		diskStatus.DiskUsed = diskStatus.DiskAll - diskStatus.DiskFree
		diskStatus.DiskAll_GB = float64(diskStatus.DiskAll) / (1024 * 1024 * 1024)
		diskStatus.DiskFree_GB = float64(diskStatus.DiskFree) / (1024 * 1024 * 1024)
		diskStatus.DiskUsed_GB = float64(diskStatus.DiskUsed) / (1024 * 1024 * 1024)
		return diskStatus, nil
	}
	return nil, errors.New("windows only,this system type not match")
}

func GetMemStatus_win() (*memoryStatusEx,error){
	GlobalMemoryStatusEx := handle.MustFindProc("GlobalMemoryStatusEx")
	var memInfo memoryStatusEx
	memInfo.dwLength = uint32(unsafe.Sizeof(memInfo))
	mem, _, err := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	if mem == 0 {
		return nil, errors.New("get windows memory failed:" + err.Error())
	}
	return &memInfo, nil
}