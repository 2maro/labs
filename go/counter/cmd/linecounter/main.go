package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := o
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines++
	}
	fmt.Println(lines)
}
