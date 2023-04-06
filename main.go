/*
tokenize-split is a command-line tool that tokenizes input text
and splits it into multiple files, with each file containing
a maximum number of tokens specified by the user.

Usage:
	$ cat too-long-text.txt | tokenize-split -n 4096
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkoukk/tiktoken-go"
)

func main() {
	// parse command-line flags
	numTokens := flag.Int("n", 4096, "maximum number of tokens per file")
	flag.Parse()

	// read input text from stdin
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
		os.Exit(1)
	}

	// initialize tiktoken-go
	encoding := "cl100k_base"
	enc, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return
	}

	// tokenize input text
	tokens := enc.Encode(string(input), nil, nil)

	// split tokens into chunks
	chunkSize := *numTokens
	numChunks := (len(tokens) + chunkSize - 1) / chunkSize // round up
	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > len(tokens) {
			end = len(tokens)
		}

		// write chunk to file
		filename := fmt.Sprintf("tokenized_%d.txt", i)
		chunk := tokens[start:end]
		text := enc.Decode(chunk)
		err := ioutil.WriteFile(filename, []byte(text), 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error writing chunk to file:", err)
			os.Exit(1)
		}
	}
}
