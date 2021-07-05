package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var list []string

	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Print("erro")
	}
	list = strings.Fields(line)
	A, _ := strconv.Atoi(list[0])
	B, _ := strconv.Atoi(list[1])
	C, _ := strconv.Atoi(list[2])
	D, _ := strconv.Atoi(list[3])
	SomaCD := C + D
	SomaAB := A + B
	if B > C && D > A && SomaCD > SomaAB && C > 0 && D > 0 && A%2 == 0 {
		fmt.Print("Valores aceitos\n")
	} else {
		fmt.Print("Valores nao aceitos\n")
	}

}
