package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"os"
)

func shred(path string) {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Error: '%s' could not be found.\n", path)
		os.Exit(1);
	}

	if ! info.Mode().IsRegular() {
		fmt.Printf("Error: '%s' is not a regular file. It cannot be shredded.\n", path)
		os.Exit(1);
	}

	passes := 3

	for passes > 0 {
		passes = passes - 1
		nonsense := make([]byte, info.Size())
		_, err :=  rand.Read(nonsense)
		if err != nil {
			fmt.Printf("Error: Could not generate random data\n");
			os.Exit(2)
		}
		write_err := os.WriteFile(path, nonsense, 0644)
		if write_err != nil {
			fmt.Printf("Error: Could not write random data\n");
			os.Exit(2)
		}
	}

	remove_err := os.Remove(path)
	if remove_err != nil {
		fmt.Printf("Error: Could not delete file.\n");
		os.Exit(2)
	}

}

func main() {

	if(len(os.Args) < 2) {
		fmt.Printf("You did not specify any path to shred.\n\n")
		fmt.Printf("Usage: %s $PATH_TO_SHRED\n", os.Args[0])
		os.Exit(1);
	}

	path := os.Args[1]

	shred(path)

}
