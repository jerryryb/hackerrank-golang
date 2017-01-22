package main
import "fmt"

func main() {
    var t int
    fmt.Scanf("%d", &t)
    for t > 0{
        t--
        var n int64
        fmt.Scanf("%d", &n)
        
        a := (n-1) % 3 
        a = n - 1 - a 
        a = a / 3
        
        b := (n-1) % 5 
        b = n - 1 - b 
        b = b / 5
        
        d := (n-1) % 15 
        d = n - 1 - d 
        d = d/15; 
        var sum int64 = 3*a*(a+1)/2 + 5*b*(b+1)/2 - 15*d*(d+1)/2
        fmt.Printf("%d\n", sum)
    }
}