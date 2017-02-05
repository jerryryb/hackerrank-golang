package main
import ("fmt"
)

func main() {
    t := 0
    fmt.Scanf("%d", &t)
    for t > 0 {
        t--
        k := 0
        nlen := 0
        fmt.Scanf("%d %d", &k, &nlen)
        s := ""
        fmt.Scanf("%s", &s)

        maxmul := 0
        for i := 0; i < len(s)-nlen; i++ {
            locmul := 1
            for j := 0; j < nlen; j++ {
                locmul *= int(s[i+j] - '0')
            }
            if locmul > maxmul {
                maxmul = locmul
            }
        }
        fmt.Printf("%d\n", maxmul)
    }
}