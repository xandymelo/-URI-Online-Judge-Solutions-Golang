package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contador := 0
	in := bufio.NewReader(os.Stdin)

	for i := 0; i < 6; i++ {
		line1, _ := in.ReadString('\n')
		line1 = strings.TrimSuffix(line1, "\n")
		line1 = strings.TrimSuffix(line1, "\r")
		line2, _ := strconv.Atoi(line1)
		if line2 > 0 {
			contador = contador + 1
		}
	}
	fmt.Printf("%d valores positivos\n", contador)
}
