package filesystem

import (
	"syscall"
	"unsafe"
	"strconv"
)
type Handle uintptr

const InvalidHandle Handle = Handle(^Handle(0))
const ERROR_MORE_DATA syscall.Errno = 234

// this would probably go in a common utilities rather than here

func convert_windows_string_list(winput []uint16) []string {
	var retstrings []string
	var rsindex = 0

	retstrings = append(retstrings, "")
	for i := 0; i < (len(winput) - 1); i++ {
		if winput[i] == 0 {
			if winput[i+1] == 0 {
				return retstrings
			}
			rsindex++
			retstrings = append(retstrings, "")
			continue
		}
		retstrings[rsindex] += string(rune(winput[i]))
	}
	return retstrings
}
// as would this
func convert_windows_string(winput []uint16) string {
	var retstring string
	for i := 0; i < len(winput); i++ {
		if winput[i] == 0 {
			break
		}
		retstring += string(rune(winput[i]))
	}
	return retstring
}
func getDiskSize(vol string) (size uint64, freespace uint64) {
	var mod = syscall.NewLazyDLL("kernel32.dll")
	var getDisk = mod.NewProc("GetDiskFreeSpaceExW")
	var sz uint64
	var fr uint64
	status, _, _ := getDisk.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(vol))),
		uintptr(0),
		uintptr(unsafe.Pointer(&sz)),
		uintptr(unsafe.Pointer(&fr)))
	if status == 0 {
		return 0, 0
	}
	return sz, fr
}
func getMountPoints(vol string) []string {
	var mod = syscall.NewLazyDLL("kernel32.dll")
	var getPaths = mod.NewProc("GetVolumePathNamesForVolumeNameW")
	var tmp uint32
	var objlistsize uint32 = 0x0
	var retval []string

	status, _, errno := getPaths.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(vol))),
		uintptr(unsafe.Pointer(&tmp)),
		2,
		uintptr(unsafe.Pointer(&objlistsize)))

	if status != 0 || errno != ERROR_MORE_DATA {
		// unexpected
		return retval
	}
	buf := make([]uint16, objlistsize)
	status, _, errno = getPaths.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(vol))),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(objlistsize),
		uintptr(unsafe.Pointer(&objlistsize)))
	if status == 0 {
		return retval
	}
	return convert_windows_string_list(buf)

}
func getFileSystemInfo() (interface{}, error) {
	var mod = syscall.NewLazyDLL("kernel32.dll")
	var findFirst = mod.NewProc("FindFirstVolumeW")
	var findNext = mod.NewProc("FindNextVolumeW")
	var findClose = mod.NewProc("FindVolumeClose")

	//var findhandle Handle
	buf := make([]uint16, 512)
	var sz int32 = 512
	fh, _, _ := findFirst.Call(uintptr(unsafe.Pointer(&buf[0])),
		uintptr(sz))
	var findhandle Handle = Handle(fh)
	var fileSystemInfo []interface{}
	
	if findhandle != InvalidHandle {
		more_data := true
		for more_data {
			outstring := convert_windows_string(buf)
			sz, _:= getDiskSize(outstring)
			var capacity string
			if 0 == sz {
				capacity = "Unknown"
			} else {
				capacity = strconv.FormatInt(int64(sz)/1024.0, 10)
			}
			mountpts := getMountPoints(outstring)
			var mount_name string
			if len(mountpts) > 0 {
				mount_name = mountpts[0]
			}
			iface := map[string]interface{}{
				"name": outstring,
				"kb_size": capacity,
				"mounted_on": mount_name,
			}
			fileSystemInfo = append(fileSystemInfo, iface)
			status, _, _ := findNext.Call(uintptr(fh),
				uintptr(unsafe.Pointer(&buf[0])),
				uintptr(sz))
			if 0 == status {
				more_data = false
			}
		}
		findClose.Call(fh)
	}
	

	return fileSystemInfo, nil
}
