package main
 
import (
    "fmt"
)
 
func main() {

    var A, B, MEDIA float64
    fmt.Scan(&A)
    fmt.Scan(&B)
    MEDIA = ( ( A * 3.5 ) + ( B * 7.5 ) ) / 11
    fmt.Printf("MEDIA = %.5f\n", MEDIA)

}
