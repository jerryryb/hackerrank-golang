package main
import "fmt"

func generatePrimes(n int64, primes []int64) []int64 {
	counter := primes[len(primes)-1] + 1
	for true {
		if n < primes[len(primes)-1] {
			break
		}
		prime := true
		for _, elem := range primes {
			if counter % elem == 0 {
				prime = false
				break
			}
            if elem * elem >= counter {
				break
			}
		}
		if prime {
			primes = append(primes, counter)
		}
		counter++
	}
	return primes
}

func main() {
    t := 0
    fmt.Scanf("%d", &t)
    primes := []int64{2,3}
    for t > 0 {
        t--
        n := int64(0)
        fmt.Scanf("%d", &n)
        if n > primes[len(primes)-1] {
            primes = generatePrimes(n, primes)
        }
        result := int64(0)
        for _, elem := range primes {
            if elem > n {
                break
            }
            result += elem
        }
        fmt.Println(result)
    }
}