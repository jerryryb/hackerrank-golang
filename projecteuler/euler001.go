package main
import "fmt"

func main() {
    var t int
    fmt.Scanf("%d", &t)
    for t > 0{
        t--
        var n int
        fmt.Scanf("%d", &n)
        var result int = 0
        for i := 0; i < n; i++ {
            if i % 3 == 0 || i % 5 == 0{
                result += i
            }
        }
        fmt.Printf("%d\n", result)
    }
}