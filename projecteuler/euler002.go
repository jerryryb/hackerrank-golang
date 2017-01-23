package main
import "fmt"

func main() {
    var t int
    fmt.Scanf("%d", &t)
    for t > 0 {
        t--
        var n int64
        fmt.Scanf("%d", &n)
        a := int64(0)
        b := int64(1)
        sum := int64(0)
        for true {
            a, b = b, a+b
            if b >= n {
                break
            }
            if b % 2 == 0 {
                sum += b
            }
        }
        fmt.Printf("%d\n", sum)
    }
}