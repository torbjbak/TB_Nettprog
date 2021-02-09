package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	isPrime := true
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}
	return isPrime
}

func findPrimes(n1, n2 int) []int {
	var primes []int
	for i := n1; i < n2; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
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
	for i := 0; i < nrSeg; i++ {
		print("seg ", i+1, ": ")
		fmt.Printf("%d\n", findPrimes(segments[i][0], segments[i][1]))
	}
}
