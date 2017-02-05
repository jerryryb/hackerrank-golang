package main
import "fmt"

func main() {
    t := 0
    fmt.Scanf("%d", &t)
    for t > 0 {
        t--
        n := 0
        fmt.Scanf("%d", &n)
        
        result := -1
        for a := 1; a < n; a++ {

            b := ((n-a)*(n-a)-(a*a))/(2 * (n-a))
            c := n-b-a

            if a*a+b*b == c*c && b > a && c > b {
                result = a * b * c
            }
        }
        fmt.Printf("%d\n", result)
    }
}