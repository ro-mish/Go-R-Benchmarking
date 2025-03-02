# This script is called by the benchmark.sh with benchmark_size set

cat("Running causal inference with dataset size:", benchmark_size, "\n")

# Generate data of the specified size
sim_data <- generate_data(n = benchmark_size)
data <- sim_data$data
true_effect <- sim_data$true_effect

# Time the execution
start_time <- Sys.time()

# Estimate the effect
results <- estimate_simple_ate(data)

end_time <- Sys.time()

# Print results
cat("Estimated ATE:", results$ate, "\n")
cat("True effect:", true_effect, "\n")
cat("Execution time:", difftime(end_time, start_time, units = "secs"), "seconds\n") 