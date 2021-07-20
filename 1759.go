package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var hoho string
	in := bufio.NewReader(os.Stdin)
	qtd_ho, _ := in.ReadString('\n')
	qtd_ho = strings.TrimSuffix(qtd_ho, "\n")
	qtd_ho = strings.TrimSuffix(qtd_ho, "\r")
	qtd_ho_int, _ := strconv.Atoi(qtd_ho)
	hoho = strings.Repeat("Ho ", qtd_ho_int)
	hoho_oficial := hoho[0:len(hoho)-1] + "!\n"
	fmt.Print(hoho_oficial)

}
