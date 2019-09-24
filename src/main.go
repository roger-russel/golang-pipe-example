package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("No pipe data was received")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	count := 0

	for {

		input, _, err := reader.ReadLine()

		if err != nil && err == io.EOF {
			break
		}

		count++

		fmt.Printf("%v\t%s\n", count, input)
	}

}
