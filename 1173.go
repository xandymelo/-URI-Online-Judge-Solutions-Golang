package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	line1, _ := in.ReadString('\n')
	line1 = strings.TrimSuffix(line1, "\n")
	line1 = strings.TrimSuffix(line1, "\r")
	line2, _ := strconv.Atoi(line1)
	for i := 0; i < 10; i++ {
		fmt.Printf("N[%d] = %d\n", i, line2)
		line2 = line2 * 2
	}

}
