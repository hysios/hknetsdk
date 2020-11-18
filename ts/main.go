package main

import (
	"io"
	"log"
	"os"
)

func main() {
	src, err := os.Open("output.mp4")
	if err != nil {
		log.Fatalf("open file %s error %s\n", "output.mp4", err)
	}

	dest, err := os.Create("output.ts")
	if err != nil {
		log.Fatalf("create file %s error %s\n", "output.ts", err)
	}

	src.Seek(40, io.SeekStart)
	log.Println(io.Copy(dest, src))
}
