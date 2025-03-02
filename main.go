package main

import (
	"flag"
	"fmt"
	"time"

	"causalinference/causalinference"
)

func main() {
	// Parse command line flags
	size := flag.Int("size", 10000, "Size of dataset to generate")
	flag.Parse()

	fmt.Printf("Running causal inference with dataset size: %d\n", *size)

	// Generate data
	start := time.Now()
	data := causalinference.GenerateCausalData(*size, 123)

	// Estimate effect
	effect := causalinference.EstimateCausalEffect(data)
	elapsed := time.Since(start)

	// Print results
	fmt.Printf("Estimated effect: %.4f\n", effect)
	fmt.Printf("True effect: %.4f\n", data.TrueEffect)
	fmt.Printf("Execution time: %.4f seconds\n", elapsed.Seconds())
}
