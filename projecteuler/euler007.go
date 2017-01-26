package main
import "fmt"

func main() {
    t := 0
    fmt.Scanf("%d", &t)
    primes := []int{2,3}
    for t > 0 {
        t--
        n := 0
        
        fmt.Scanf("%d", &n)
        if len(primes) >= n {
            fmt.Printf("%d\n", primes[n-1])
        }else {
            counter := primes[len(primes)-1] + 1
            for len(primes) <= n {
                if counter % 2 != 0 && counter % 3 != 0 {
                    temp := 4
                    for temp * temp <= counter {
                        if counter % temp == 0 {
                            break
                        }
                        temp++
                    }
                    if temp * temp > counter {
                        primes = append(primes, counter)
                    }
                }
                counter++
            }
            fmt.Printf("%d\n", primes[n-1])
        }
    }
}