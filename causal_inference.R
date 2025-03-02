# Minimal causal inference implementation

# Generate synthetic data
generate_data <- function(n = 1000, seed = 42) {
  set.seed(seed)

  # Create single covariate X
  X <- rnorm(n)

  # Treatment depends on X
  prob_treat <- 0.5 * (X + 1)
  treatment <- as.integer(runif(n) < prob_treat)

  # True effect
  true_effect <- 5.0

  # Outcome depends on X and treatment
  outcome <- X + true_effect * treatment + rnorm(n)

  # Return as a simple data frame
  data <- data.frame(
    X = X,
    treatment = treatment,
    outcome = outcome
  )

  return(list(data = data, true_effect = true_effect))
}

# Estimates simple treatment effect
estimate_simple_ate <- function(data) {
  start_time <- Sys.time()

  # Simple difference in means
  treated <- data$outcome[data$treatment == 1]
  control <- data$outcome[data$treatment == 0]

  ate <- mean(treated) - mean(control)

  end_time <- Sys.time()

  return(list(
    ate = ate,
    execution_time = difftime(end_time, start_time, units = "secs")
  ))
}