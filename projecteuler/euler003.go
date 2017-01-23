package main
import "fmt"

func largest_prime_factor(n int64) (largest_factor int64) {
    largest_factor = int64(1)
 
    // remove any factors of 2 first
    for n % 2 == 0 {
        largest_factor = 2
        n /= 2
    }
    
    // now look at odd factors
    p := int64(3)
    for n != 1 {
        for n % p == 0 {
            largest_factor = p
            n /= p
        }
        p += 2
    }
    return largest_factor
}

func main() {
    var t int
    fmt.Scanf("%d", &t)
    for t > 0{
        t--
        var n int64
        fmt.Scanf("%d", &n)
        fmt.Println(largest_prime_factor(n))
    }
}