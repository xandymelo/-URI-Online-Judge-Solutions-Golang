package main
 
import (
    "fmt"
)
 
func main() {

    var A, B, C, D, DIF int
    fmt.Scan(&A)
    fmt.Scan(&B)
    fmt.Scan(&C)
    fmt.Scan(&D)
    DIF = ( ( A * B ) - ( C * D ) )
    fmt.Printf("DIFERENCA = %d\n", DIF)

}
