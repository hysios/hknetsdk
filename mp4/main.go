package main

import (
	"fmt"
	"os"

	"github.com/alfg/mp4"
)

func main() {
	file, err := os.Open("output.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	size := info.Size()

	mp4, _ := mp4.OpenFromReader(file, size)
	file.Close()
	fmt.Println(file.Ftyp.Name)
	fmt.Println(file.Ftyp.MajorBrand)
}
