package main

import "fmt"

type ByteSize float64

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {

	fmt.Println("======== Conversion Chart ========")
	fmt.Println("1KB =", KB, "bytes")
	fmt.Println("1MB =", MB, "bytes")
	fmt.Println("1GB =", GB, "bytes")
	fmt.Println("1TB =", TB, "bytes")
	fmt.Println("1PB =", PB, "bytes")
	fmt.Println("1EB =", EB, "bytes")
	fmt.Println("1ZB =", float64(ZB), "bytes")
	fmt.Println("1YB =", float64(YB), "bytes")
	fmt.Println("==================================")

	fmt.Println("Some tests:")
	fmt.Println(ByteSize(YB))
	fmt.Println(ByteSize(1000))
	fmt.Println(ByteSize(5000))
	fmt.Println(ByteSize(7500000))
	fmt.Println(ByteSize(8926711598))
	fmt.Println(ByteSize(1e13))
}

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}
