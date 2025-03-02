#!/bin/bash

# cmd to compile Go code
go build -o causal_inference_go

# Run benchmarks with different sizes
for size in 1000 10000 100000; do
    echo "Size: $size"
    
    # Go benchmarking test
    go_start=$(date +%s.%N)
    ./causal_inference_go -size=$size

    go_end=$(date +%s.%N)
    go_time=$(echo "$go_end - $go_start" | bc)

    # R benchmark
    r_start=$(date +%s.%N)
    Rscript -e "source('causal_inference.R'); benchmark_size <- $size; source('run_benchmark.R')"

    r_end=$(date +%s.%N)
    r_time=$(echo "$r_end - $r_start" | bc)
    
    # Show results
    echo "Go: $go_time s, R: $r_time s"
done 