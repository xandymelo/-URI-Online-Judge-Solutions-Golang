package main
 
import (
    "fmt"
)
 
func main() {

    var A, B, C, MEDIA float64
    fmt.Scan(&A)
    fmt.Scan(&B)
    fmt.Scan(&C)
    MEDIA = ( ( A * 2 ) + ( B * 3 ) + ( C * 5 ) ) / 10
    fmt.Printf("MEDIA = %.1f\n", MEDIA)

}
