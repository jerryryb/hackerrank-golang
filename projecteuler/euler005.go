package main
import "fmt"

func main() {
    t := 0
    fmt.Scanf("%d", &t)
    for t > 0 {
        t--
        n := 0
        fmt.Scanf("%d", &n)
        
        x := 1
        for {
            x++
            result := true
            for i := 1; i <= n; i++ {
                if x % i != 0 {
                    result = false
                    break
                }                
            }
            if result {
                break
            }
        }
        fmt.Printf("%d\n", x)
    }
}