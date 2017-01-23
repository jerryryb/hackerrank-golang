package main
import ("fmt"
        _"math"
       )

func main() {
    var t int
    fmt.Scanf("%d", &t)
    for t > 0{
        t--
        var n int64
        fmt.Scanf("%d", &n)

        i := int64(2)
        for i * i < n {
            for n % i == 0 {
                n = n / i
            }
            i = i + 1
        }
        fmt.Println(n)
    }
}