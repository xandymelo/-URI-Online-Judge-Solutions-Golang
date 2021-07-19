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
	var posicao int
	var valor_atual float64

	menorvalor := math.Inf(2)
	in := bufio.NewReader(os.Stdin)
	tamanho_vetor, _ := in.ReadString('\n')
	tamanho_vetor = strings.TrimSuffix(tamanho_vetor, "\n")
	tamanho_vetor = strings.TrimSuffix(tamanho_vetor, "\r")
	line2, _ := strconv.Atoi(tamanho_vetor)
	elementos, _ := in.ReadString('\n')
	elementos = strings.TrimSuffix(elementos, "\n")
	elementos = strings.TrimSuffix(elementos, "\r")
	lista_elementos := strings.Split(elementos, " ")
	for i := 0; i < line2; i++ {
		valor_atual, _ = strconv.ParseFloat(lista_elementos[i], 64)
		if menorvalor > valor_atual {
			menorvalor = valor_atual
			posicao = i
		}
	}
	menorvalorint := int(menorvalor)
	fmt.Printf("Menor valor: %d\n", menorvalorint)
	fmt.Printf("Posicao: %d\n", posicao)

}
