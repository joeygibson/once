package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var exists = struct{}{}
	reader := bufio.NewReader(os.Stdin)
	seen := make(map[string]struct{})

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if _, ok := seen[line]; ok {
			continue
		}

		seen[line] = exists

		fmt.Print(line)
	}
}
