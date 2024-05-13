package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
)

func main() {
	flag.Parse()

	filename := flag.Arg(0)
	if filename == "" {
		filename = os.Stdin.Name()
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("there was an error reading file:", err)
		os.Exit(1)
	}

	words := bytes.Split(data, []byte("\n"))
	// remove empty space on last split
	words = words[:len(words)-1]

	sort.Slice(words, func(i, j int) bool {
		return bytes.Compare(words[i], words[j]) < 0
	})

	for _, word := range words {
		os.Stdout.Write(word)         // Write the byte slice
		os.Stdout.Write([]byte("\n")) // Write a newline
	}
}
