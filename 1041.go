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
	x, _ := strconv.ParseFloat(list[0], 64)
	y, _ := strconv.ParseFloat(list[1], 64)

	if x > 0 && y > 0 {
		fmt.Print("Q1\n")
	} else if x < 0 && y > 0 {
		fmt.Print("Q2\n")
	} else if x < 0 && y < 0 {
		fmt.Print("Q3\n")
	} else if x > 0 && y < 0 {
		fmt.Print("Q4\n")
	} else if x == 0 && y == 0 {
		fmt.Print("Origem\n")
	} else if x == 0 {
		fmt.Print("Eixo Y\n")
	} else if y == 0 {
		fmt.Print("Eixo X\n")
	}

}
