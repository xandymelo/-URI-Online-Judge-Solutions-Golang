package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var vetor []int
	in := bufio.NewReader(os.Stdin)

	for i := 0; i < 11; i++ {
		line1, _ := in.ReadString('\n')
		line1 = strings.TrimSuffix(line1, "\n")
		line1 = strings.TrimSuffix(line1, "\r")
		line2, _ := strconv.Atoi(line1)
		if line2 <= 0 {
			line2 = 1
		}
		vetor = append(vetor, line2)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("X[%d] = %d\n", i, vetor[i])
	}

}
