package util

import "strconv"

func StringPtrUint(x string) (y uint) {
	re, _ := strconv.Atoi(x)
	y = uint(re)
	return
}

func StringPtrInt(x string) (y int) {
	y, _ = strconv.Atoi(x)
	return
}

func StringPtrInt64(x string) (y int64) {
	y, _ = strconv.ParseInt(x, 10, 64)
	return
}

func UintPtrString(x uint) (y string) {
	y = strconv.Itoa(int(x))
	return
}

func Int64PtrString(x int64) (y string) {
	y = strconv.Itoa(int(x))
	return
}
