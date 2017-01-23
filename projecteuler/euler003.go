package main
import ("fmt"
        "math"
)

func largest_prime_factor(n int64) (ans int64) {
    for n % 2 == 0 {
      n /= 2
    }
    
    if n == 1 {
        ans = 2
        return
    }
    
    var i int64
    for i = 3; i <= int64(math.Floor(math.Sqrt(float64(n)))); i = i + 2 {
        for n % i == 0 {
            n /= i
        }
    }
    
    if n > 2 {
        ans = n
    }else{
        ans = i - 2
    }
    return ans
}

func main() {
    var t int
    fmt.Scanf("%d", &t)
    for t > 0 {
        t--
        var n int64
        fmt.Scanf("%d", &n)
        fmt.Println(largest_prime_factor(n))
    }
}