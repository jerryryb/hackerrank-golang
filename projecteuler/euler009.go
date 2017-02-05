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
        for a := 1; a < n-2; a++ {
            for b := a+1; b < n-1; b++ {
                if a + b > n {
                    break
                }
                for c := b+1; c < n; c++ {
                    if a + b + c > n {
                        break
                    }
                    
                    if a + b + c == n {
                        if a*a + b*b == c*c {
                            result = a * b * c
                        }
                    }
                }
            }
        }
        fmt.Printf("%d\n", result)
    }
}