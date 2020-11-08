// +build linux darwin

package Utils

import "errors"

//内存空间使用情况
type MemStatus struct {
	MemAll  uint32 `json:"all"`
	MemUsed uint32 `json:"used"`
	MemFree uint32 `json:"free"`
	MemSelf uint64 `json:"self"`
}


//磁盘空间使用情况
type DistStatus struct{
	DiskAll		int64
	DiskFree	int64
	DiskUsed	int64
	DiskAll_GB	int
	DiskFree_GB	int
	DiskUsed_GB	int
}
 /******
 获取linux系统下给定路径的磁盘空间情况
 ******/
func GetDiskSpaceStatus_linux(path string){
	sysType := runtime.GOOS
	if sysType == "windows" {
		if len(path) <= 0 {
			path = "/"
		}
		fs := syscall.Statfs_t
		err := syscall.Statfs(path, &fs)
		if err != nil {
			return nil, err
		}

		diskStatus := new(DistStatus)
		diskStatus.DiskAll = fs.Blocks * uint64(fs.Bsize)
		diskStatus.DiskFree = fs.Bfree * uint64(fs.Bsize)
		diskStatus.DiskUsed = diskStatus.DiskAll - diskStatus.DiskFree
		diskStatus.DiskAll_GB = float64(diskStatus.DiskAll) / (1024 * 1024 * 1024)
		diskStatus.DiskFree_GB = float64(diskStatus.DiskFree) / (1024 * 1024 * 1024)
		diskStatus.DiskUsed_GB = float64(diskStatus.DiskUsed) / (1024 * 1024 * 1024)
		return diskStatus, nil
	}
	return nil, errors.New("linux only,this system type not match")
}

func GetMemStatus_linux()(*MemStatus, error){
	//查询程序自身内存占用情况
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	mem := new(MemStatus)
	mem.MemSelf = memStat.Allo

	sysInfo := new(syscall.Sysinfo_t)
	err := syscall.Sysinfo(sysInfo)
	if err != nil {
		return nil, errors.New("get linux memory space status failed:" + err.Error())
	}
	mem.MemAll = sysInfo.Totalram * uint32(syscall.Getpagesize())
	mem.MemFree = sysInfo.Freeram * uint32(syscall.Getpagesize())
	mem.MemUsed = mem.All - mem.Free
	return mem,nil
}