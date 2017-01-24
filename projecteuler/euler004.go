package main
import ("fmt"
        "strconv"
)

func ispalindrom(n int) (result bool) {
    s := strconv.Itoa(n)
    length := len(s)
    for i := 0; i < length/2; i++ {
        if s[i] != s[length-i-1] {
            return false
        }
    }
    return true
}

func main() {
    var t int
    fmt.Scanf("%d", &t)
    for t > 0 {
        t--
        var n int
        fmt.Scanf("%d", &n)
        var max_palindrom int = 0
        for i := 100; i < 1000; i++ {
            for j := 100; j < 1000; j++ {
                var res int = i*j
                if res >= n {
                    break
                }
                if res > max_palindrom && ispalindrom(res) {
                    max_palindrom = res
                }
            }
        }
        fmt.Printf("%d\n", max_palindrom)
    }
}