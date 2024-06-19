package main

import (
	"fmt"
	"strconv"
)

func getFilePermissions(flag int) string {
	v := int64(flag)
	c := strconv.FormatInt(v, 10)
	f, err := strconv.ParseInt(c, 8, 64)
	if err != nil {
		panic(err)
	}
	OwnerR := strconv.FormatInt(f&0400>>8, 2)
	OwnerW := strconv.FormatInt(f&0200>>7, 2)
	OwnerE := strconv.FormatInt(f&0100>>6, 2)

	GroupR := strconv.FormatInt(f&040>>5, 2)
	GroupW := strconv.FormatInt(f&020>>4, 2)
	GroupE := strconv.FormatInt(f&010>>3, 2)

	OtherR := strconv.FormatInt(f&04>>2, 2)
	OtherW := strconv.FormatInt(f&02>>1, 2)
	OtherE := strconv.FormatInt(f&01, 2)

	OwR := checkRead(OwnerR)
	GrR := checkRead(GroupR)
	OtR := checkRead(OtherR)

	OwW := checkWrite(OwnerW)
	GrW := checkWrite(GroupW)
	OtW := checkWrite(OtherW)

	OwE := checkExecute(OwnerE)
	GrE := checkExecute(GroupE)
	OtE := checkExecute(OtherE)

	return fmt.Sprintf("Owner:%v,%v,%v Group:%v,%v,%v Other:%v,%v,%v",
		OwR, OwW, OwE, GrR, GrW, GrE, OtR, OtW, OtE)
}

func checkRead(num string) string {
	if num != "0" {
		return "Read"
	}
	return "-"
}
func checkWrite(num string) string {
	if num != "0" {
		return "Write"
	}
	return "-"
}

func checkExecute(num string) string {
	if num != "0" {
		return "Execute"
	}
	return "-"
}

func main() {
	fmt.Println(getFilePermissions(777))
	var s uint8
	s = 1 << 7
	fmt.Println(s)
}
