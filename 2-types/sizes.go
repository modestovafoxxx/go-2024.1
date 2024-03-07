package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("sizeof bool = ", unsafe.Sizeof(bool(false)))

	fmt.Println("sizeof string = ", unsafe.Sizeof(""))
	fmt.Println("sizeof long string = ", unsafe.Sizeof("IyACQk3Y3b2JDPA8a1W2qIMKBvV9snyvCNdu70qBkHobmgT2g2VwUNHUX0fHKSZ46FK96vv4NTEdW8JfGyFuAx1tJgBT6B0ZPJFO"))

	fmt.Println("sizeof int = ", unsafe.Sizeof(int(0)))
	fmt.Println("sizeof int8 = ", unsafe.Sizeof(int8(0)))
	fmt.Println("sizeof int16 = ", unsafe.Sizeof(int16(0)))
	fmt.Println("sizeof int32 = ", unsafe.Sizeof(int32(0)))
	fmt.Println("sizeof int64 = ", unsafe.Sizeof(int64(0)))

	fmt.Println("sizeof uint = ", unsafe.Sizeof(uint(0)))
	fmt.Println("sizeof uint8 = ", unsafe.Sizeof(uint8(0)))
	fmt.Println("sizeof uint16 = ", unsafe.Sizeof(uint16(0)))
	fmt.Println("sizeof uint32 = ", unsafe.Sizeof(uint32(0)))
	fmt.Println("sizeof uint64 = ", unsafe.Sizeof(uint64(0)))

	fmt.Println("sizeof uintptr = ", unsafe.Sizeof(uintptr(0)))

	fmt.Println("sizeof byte = ", unsafe.Sizeof(byte(0)))

	fmt.Println("sizeof rune = ", unsafe.Sizeof(rune(0)))

	fmt.Println("sizeof float32 = ", unsafe.Sizeof(float32(0)))
	fmt.Println("sizeof float64 = ", unsafe.Sizeof(float64(0)))

	fmt.Println("sizeof complex64 = ", unsafe.Sizeof(complex64(0)))
	fmt.Println("sizeof complex128 = ", unsafe.Sizeof(complex128(0)))

	fmt.Println("sizeof slice = ", unsafe.Sizeof([]int{}))
	fmt.Println("sizeof init slice = ", unsafe.Sizeof([]int{1, 2, 3}))
}
