package basics

import (
	"fmt"
)

func ConstantOPs() {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Constant OPs")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	const (
		n  int     = 1
		ij float32 = 78.
	)

	fmt.Printf("%v , %T\n", n, n)
	fmt.Printf("%v , %T\n", ij, ij)

	fmt.Println("iota OPs")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
	const (
		i = iota
		j = iota
	)
	fmt.Printf("%v , %T\n", i, i)
	fmt.Printf("%v , %T\n", j, j)
	fmt.Println("iota BitWise OPs")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")

	// can be used as enum
	const (
		doctor = iota
		nurse  = iota
		helper = iota
	)
	fmt.Printf(" Doctor %v , %T\n", doctor, doctor)
	fmt.Printf(" Nurse  %v , %T\n", nurse, nurse)
	fmt.Printf(" Helper %v , %T\n", helper, helper)

	const (
		_ = iota //blank identifier
		// above operation is to ignore zero value
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

	fileSize := 1024.0 * 1024. * 1024.
	fmt.Printf("File size in KB %.2f KB \n", fileSize/KB)
	fmt.Printf("File size in MB %.2f MB\n", fileSize/MB)
	fmt.Printf("File size in GB %.2f GB\n", fileSize/GB)
	fmt.Printf("File size in PB %.2f PB\n", fileSize/PB)

}
