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
	var preco, valor_total float64
	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Print("erro")
	}
	list = strings.Fields(line)
	codigo_item, _ := strconv.Atoi(list[0])
	qtd_item, _ := strconv.ParseFloat(list[1], 64)
	switch codigo_item {
	case 1:
		preco = 4
	case 2:
		preco = 4.5
	case 3:
		preco = 5
	case 4:
		preco = 2
	case 5:
		preco = 1.5
	}
	valor_total = qtd_item * preco
	fmt.Printf("Total: R$ %.2f\n", valor_total)
}
