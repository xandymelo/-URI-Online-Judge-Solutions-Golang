package main
 
import (
    "fmt"
)
 
func main() {

    var n, raio, area float64
    n = 3.14159
    fmt.Scan(&raio)
    area = n * ( raio * raio )
    fmt.Printf("A=%.4f\n", area)

}
