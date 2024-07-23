package main

import (
	"fmt"
	"os"

	gornnoise "github.com/shenjinti/go-rnnoise"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a raw/pcm file")
		fmt.Println("Usage: gornnoise <raw/pcm file> <output file>")
		os.Exit(1)
	}
	args := os.Args[1:]
	inFile, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	outFile, err := os.Create(args[1])
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	rnnoise := gornnoise.NewRNNoise()
	frame := make([]byte, gornnoise.GetFrameSize()*2)
	fmt.Println("Input file:", args[0], "frame size:", len(frame))
	for {
		n, err := inFile.Read(frame)
		if err != nil {
			break
		}
		if n != len(frame) {
			break
		}
		out, _ := rnnoise.Process(frame)
		outFile.Write(out)
	}
	fmt.Println("Output file:", args[1])
}
