package main

import (
	"bufio"
	"fmt"
	"math"
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
	A, _ := strconv.ParseFloat(list[0], 64)
	B, _ := strconv.ParseFloat(list[1], 64)
	C, _ := strconv.ParseFloat(list[2], 64)

	delta := (B * B) - (4 * A * C)

	if delta < 0 || A == 0.0 {
		fmt.Print("Impossivel calcular\n")
	} else {
		X1 := (-B + math.Sqrt(delta)) / (2 * A)
		X2 := (-B - math.Sqrt(delta)) / (2 * A)
		fmt.Printf("R1 = %.5f\n", X1)

		fmt.Printf("R2 = %.5f\n", X2)
	}

}
