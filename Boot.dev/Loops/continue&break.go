package main

import "fmt"

func printPrimes(max int) {
    for i:=2; i<=max; i++ {
        count := 0
        for j:=1; j<=i; j++ {
            if i%j == 0 {
                count++
            }
            if count == 3 {
                break
            }
        }
        if count < 3{
            fmt.Println(i)
        }
    }
}

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
