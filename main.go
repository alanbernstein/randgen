package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var distribution string
	var max, count int
	var seed int64
	var exponent, offset float64
	flag.Int64Var(&seed, "seed", -1, "PRNG seed; omit for time-based")
	flag.StringVar(&distribution, "dist", "zipf", "distribution")
	flag.Float64Var(&exponent, "exponent", 1.2, "exponent for zipf")
	flag.Float64Var(&offset, "offset", 1.1, "offset for zipf")
	flag.IntVar(&max, "max", 100, "maximum for [zipf, uniform]")
	flag.IntVar(&count, "count", 1000, "count (number of samples)")
	flag.Parse()

	if seed == -1 {
		// rand.Seed(time.Now().UTC().UnixNano())
		seed = time.Now().UTC().UnixNano()
	}

	switch distribution {
	case "zipf":
		genZipf(seed, exponent, offset, max, count)
	case "uniform":
		genUniform(seed, 0, 100, count)
	default:

	}

}

func genUniform(seed int64, min, max, numSamples int) {
	for n := 0; n < numSamples; n++ {
		fmt.Println("0")
	}
}

func genZipf(seed int64, exponent, offset float64, max, numSamples int) {
	rnd := rand.New(rand.NewSource(seed))
	rng := rand.NewZipf(rnd, exponent, offset, uint64(max))
	for n := 0; n < numSamples; n++ {
		fmt.Println(rng.Uint64())
	}
}
