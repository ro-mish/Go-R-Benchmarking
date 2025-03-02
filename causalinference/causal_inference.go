package causalinference

import "math/rand"

// CausalData struct for building synthetic data objects
type CausalData struct {
	X          []float64 // single covariate
	Treatment  []int     // 0 or 1
	Outcome    []float64 // observed outcome
	TrueEffect float64   // for testing
}

// GenerateCausalData creates synthetic data
func GenerateCausalData(n int, seed int64) *CausalData {
	rand.Seed(seed)

	data := &CausalData{
		X:          make([]float64, n),
		Treatment:  make([]int, n),
		Outcome:    make([]float64, n),
		TrueEffect: 5.0,
	}

	for i := 0; i < n; i++ {
		// Generate basic data
		data.X[i] = rand.NormFloat64()

		// Treatment is more likely for higher X values
		if rand.Float64() < 0.5*(data.X[i]+1) {
			data.Treatment[i] = 1
		}

		// Outcome depends on X and treatment
		data.Outcome[i] = data.X[i] + float64(data.Treatment[i])*data.TrueEffect + rand.NormFloat64()
	}

	return data
}

// EstimateCausalEffect checks difference in means between treatment and control groups
func EstimateCausalEffect(data *CausalData) float64 {
	var treatSum, controlSum float64
	var treatCount, controlCount int

	// Simple means by treatment group
	for i := range data.X {
		if data.Treatment[i] == 1 {
			treatSum += data.Outcome[i]
			treatCount++
		} else {
			controlSum += data.Outcome[i]
			controlCount++
		}
	}

	// Edge case where no treatment or control observations
	if treatCount == 0 || controlCount == 0 {
		return 0
	}

	return (treatSum / float64(treatCount)) - (controlSum / float64(controlCount))
}
