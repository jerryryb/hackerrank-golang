package main
import "fmt"
func abs(n int)(int){
    if n >= 0 {
        return n
    }
    return -n
}
func main() {
    t := 0
    fmt.Scanf("%d", &t)
    for t > 0 {
        t--
        n := 0
        fmt.Scanf("%d", &n)
        
        a := 0
        b := 0
        for i := 1; i <= n; i++ {
            a += i
            b += i*i
        }
        a *= a
        
        fmt.Printf("%d\n", abs(a - b))
    }
}