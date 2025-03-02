package causalinference

import "testing"

func TestBasicFunctionality(t *testing.T) {
	// Generate a tiny dataset and verify basic properties
	data := GenerateCausalData(20, 123)

	// If data not created correctly, throw an error
	if len(data.X) != 20 || len(data.Treatment) != 20 || len(data.Outcome) != 20 {
		t.Error("Data arrays have incorrect length")
	}

	// Make sure effect is not 0
	effect := EstimateCausalEffect(data)

	if effect == 0 {
		t.Error("Estimated effect should not be zero")
	}
}

func BenchmarkAll(b *testing.B) {
	// Combined benchmark for the entire workflow
	for i := 0; i < b.N; i++ {
		data := GenerateCausalData(500, int64(i))
		EstimateCausalEffect(data)
	}
}
