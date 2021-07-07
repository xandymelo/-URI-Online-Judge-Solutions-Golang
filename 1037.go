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

	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Print("erro")
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	line_float, err := strconv.ParseFloat(line, 64)
	if line_float >= 0 && line_float <= 25 {
		fmt.Print("Intervalo [0,25]\n")
	} else if line_float > 25 && line_float <= 50 {
		fmt.Print("Intervalo (25,50]\n")
	} else if line_float > 50 && line_float <= 75 {
		fmt.Print("Intervalo (50,75]\n")
	} else if line_float > 75 && line_float <= 100 {
		fmt.Print("Intervalo (75,100]\n")
	} else {
		fmt.Print("Fora de intervalo\n")
	}
}
