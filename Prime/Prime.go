package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"sync"
)

func isPrime(n int) bool {
	prime := true

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}
	return prime
}

func findPrimes(n1, n2 int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			ch <- i
		}
	}
}

func split(n1, n2, nrSeg int) [][2]int {
	var segments [][2]int
	size := (n2 - n1 + 1) / nrSeg
	rest := (n2 - n1 + 1) % nrSeg
	counter := n1

	for i := 1; i <= nrSeg; i++ {
		if i <= rest {
			segments = append(segments, [2]int{counter, counter + size})
			counter = counter + size + 1
		} else {
			segments = append(segments, [2]int{counter, counter + size - 1})
			counter = counter + size
		}
	}
	return segments
}

func main() {
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])
	nrSeg, _ := strconv.Atoi(os.Args[3])
	segments := split(n1, n2, nrSeg)
	var wg sync.WaitGroup
	ch := make(chan int)

	fmt.Println(segments)
	for i := 0; i < nrSeg; i++ {
		wg.Add(1)
		go findPrimes(segments[i][0], segments[i][1], &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var primes []int
	for k := range ch {
		primes = append(primes, k)
	}

	sort.Ints(primes)
	fmt.Println(primes)
}
