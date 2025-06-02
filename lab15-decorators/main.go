package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"time"
)

type piFunc func(int) float64

func wraplogger(fun piFunc, logger *log.Logger) piFunc {
	return func(n int) float64 {
		fn := func(n int) (result float64) { // result gives us access to the return value of the closure (the return fun(n)) part
			defer func(t time.Time) {
				logger.Printf("took=%v, n=%v, result=%v", time.Since(t), n, result)
			}(time.Now())
			return fun(n)
		}
		return fn(n)
	}
}

func wrapcache(fun piFunc, cache *sync.Map) piFunc { // A map primitive that is thread safe!

	return func(n int) float64 {
		fn := func(n int) float64 {
			key := fmt.Sprintf("n=%d", n)
			val, ok := cache.Load(key)

			if ok {
				return val.(float64)
			}

			result := fun(n)
			cache.Store(key, result)
			return result

		}
		return fn(n)
	}

}

func Pi(n int) float64 {
	ch := make(chan float64)

	for k := 0; k <= n; k++ {
		go func(ch chan float64, k float64) {
			ch <- 4 * math.Pow(-1, k) / (2*k + 1)
		}(ch, float64(k))
	}

	result := 0.0

	for k := 0; k <= n; k++ {
		result += <-ch
	}

	return result
}

func divideByTwo(n int) float64 {
	return float64(n / 2)
}

func main() {
	// fmt.Println(Pi(1000))
	// fmt.Println(Pi(50000))

	// f := wraplogger(Pi, log.New(os.Stdout, "test ", 1))
	// f(1000)
	// f(50000)
	// f(100000)

	f := wrapcache(Pi, &sync.Map{})
	g := wraplogger(f, log.New(os.Stdout, "test ", 1))

	g(100_000)
	g(20_000)
	g(100_000)

	h := wrapcache(divideByTwo, &sync.Map{})
	i := wraplogger(h, log.New(os.Stdout, "test ", 1))

	i(100_000)
	i(20_000)
	i(100_000)

}
