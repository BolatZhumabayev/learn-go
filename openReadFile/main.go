package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("write fileName")
		os.Exit(1)
	}
	fmt.Println(os.Args[1])
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File not opened, error", err)
		os.Exit(1)
	}

	defer f.Close()
	io.Copy(os.Stdout, f)
}
