# Background
## What is Causal Inference?
To understand causal inference, we start by examining the relationship between two variables: a predictor variable (A) and a dependent variable (B). Our first goal is to establish correlation, which indicates that an increase in the dependent variable B is likely associated with an increase in the predictor variable A. However, establishing causation is a different objective altogether. 

To determine causation, we utilize various techniques employed by scientists and researchers across disciplines such as economics, epidemiology, and education. These methods help us assess whether one variable actually causes an increase in another. 

For instance, in the healthcare industry, ethical constraints often prevent certain experiments from being conducted on humans. Consequently, we may rely on simulations and hypotheses based on pre-existing data to explore the causal links between various factors and their outcomes. 

By applying techniques such as comparing the means of treatment groups and analyzing the efficacy rates of different interventions, we can gain deeper insights into whether certain factors are indeed causal for others.

## How do we implement this in our programs?
To implement our programs, we plan to use R and Go. R offers several packages that enable us to deeply examine causal inference. Specifically, for this experiment, we are focusing on propensity score matching, which allows us to compare differences in means across various effect sizes and assess the propensity scores for specific variables. Utilizing these packages in R is important because many experiments are already established in this language, which leads to faster development times by reusing existing libraries. Since there is no unified framework for developing these types of experiments, R and its pre-existing libraries provide a flexible approach that allows for a top-down programming style to run different experiments.

On the other hand, for Go, we are developing many of these techniques from the ground up. Although Go is not typically the language statisticians use for these programs, it enables us to design core functions using libraries such as the Random and Math libraries, helping us build that essential functionality.

## R Implementation Walkthrough
In the generate data function, we take several parameters as inputs. The first parameter, **n**, represents the number of examples. We also set a random seed of **42** to ensure that the code runs consistently and produces the same random results. 

We start by creating a covariate matrix **X**, which is generated from a normal distribution using the **rnorm** function in R. Next, we create the probabilities of treatment, the treatment variables, and set the true effect size to **5.0**. The outcome is influenced by both **X** and the treatment, resulting in a dataset that we can work with.
```
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
```
After that, we create a function called `estimate_simple_ate`, which estimates the treatment effect. This process begins by recording the system time, and then we analyze the difference in means between the treatment groups where the outcome is one and the control groups where the outcome is zero. These groups represent the treated and controlled populations, respectively. Next, we compute the difference in means and finally record the end time. The function returns both the estimated treatment effects and the execution time.

```
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
```

## Go Implementation
In Go, we can utilize various programming concepts, such as creating a struct to store all the relevant data for the causal data structure. Additionally, we have a function called `generateCausalData`, which generates causal data in a manner similar to the R program. This function creates random examples from a normal distribution, using a seed value of 42.

Next, we estimate the causal effect by calculating the simple means for each treatment group and then computing the difference between them. We also handle edge cases, such as when both the treatment and control counts are equal to zero; in such cases, we return a value of zero. Finally, we return both differences calculated from the treatment groups.
```
// CausalData struct for building synthetic data objects
type CausalData struct {
	X          []float64 // single covariate
	Treatment  []int     // 0 or 1
	Outcome    []float64 // observed outcome
	TrueEffect float64   // for testing
}

```

## Putting it all together
To conduct these experiments, we found that using a shell script was an effective approach. This enabled us to run benchmarks with various dataset sizes and compare the performance improvements between R and Go. For the benchmarks, we recorded the start and end times, which we displayed in the command line each time we executed the script. The script utilized a simple for loop to output the different benchmark times.

## Results
```
Size: 1000
Running causal inference with dataset size: 1000
Estimated effect: 6.3741
True effect: 5.0000
Execution time: 0.0001 seconds
Running causal inference with dataset size: 1000 
Estimated ATE: 6.408732 
True effect: 5 
Execution time: 0.0002040863 seconds
Go: 0 s, R: 0 s

Size: 10000
Running causal inference with dataset size: 10000
Estimated effect: 6.4296
True effect: 5.0000
Execution time: 0.0004 seconds
Running causal inference with dataset size: 10000 
Estimated ATE: 6.333761 
True effect: 5 
Execution time: 0.0003149509 seconds
Go: 0 s, R: 0 s

Size: 100000
Running causal inference with dataset size: 100000
Estimated effect: 6.3720
True effect: 5.0000
Execution time: 0.0041 seconds
Running causal inference with dataset size: 1e+05 
Estimated ATE: 6.37864 
True effect: 5 
Execution time: 0.001379967 seconds
Go: 0 s, R: 0 s
```

In examining the performance improvements between R and Go, we find that R actually outperforms Go in several categories, particularly with larger data set dimensions. Specifically, for smaller data sets, Go tends to perform better, but as the size increases, R takes the lead. It is important to note that using multi-threaded programs could enhance performance even further. However, in our current evaluation of Go, we have not utilized any goroutines. By implementing goroutines, we could take advantage of the computer's threads, which would significantly improve processing speed by leveraging a multi-threaded approach.

-----
#GenAI Disclosure
For this project I used GenAI for a few different cases. One was to build an intuition around what counterfactual causal inference really was and the others were to just give me code examples of using R and Go for these tasks and finally, just being able to use the shell scripts for putting all the experiments together. Though, they did give me some examples, I ended up modifying almost all of it in order to match what we've been doing and kind of being able to use my knowledge of Go and translate that GenAI output towards building out the program itself. Below is the conversation I had with chat GPT. 

```
You said: Explain counterfactual causal inference and give me concrete examples and use-cases of how it works.


Reasoned about counterfactual causal inference for a few seconds
Beginner-Friendly Explanation of Counterfactual Causal Inference

Suppose you want to know if a new teaching method actually causes students to learn better. You can give some students this new teaching method and compare their test scores to students who learned the “old” way. But even if you see a difference in scores, how do you know whether the new teaching method caused the improvement, as opposed to other factors—like having more motivated students in the new-method class?

This is the basic challenge of causal inference: figuring out whether one thing causes another, rather than just being associated with it. Here’s the key step:

Correlation vs. Causation

Correlation: “A and B happen together.”
Causation: “A causes B to happen.”
We often see a pattern in data (say, more studying is linked to higher test scores), but that doesn’t prove that studying causes higher scores. Maybe students who study more also use better study guides. In causal inference, we try to separate out these different explanations.
Counterfactuals (or Potential Outcomes)

A counterfactual is a “what if?” scenario. Imagine a specific student: in reality, they either got the new teaching method or they didn’t. But for causal inference, we ask: “What if the same student had gotten the method and what if they had not?”
We can’t actually observe both situations for the same student at the same time. So we try to compare groups that are as similar as possible—except for whether or not they got the new method. If done well, that can help us isolate the method’s causal effect.
Why Different Fields Care About This

Economics: Economists use these ideas to understand policies (like how a tax change causes people to spend more or less).
Epidemiology: Scientists studying health often can’t run certain experiments (for ethical reasons), so they must rely on real-world data and methods to see if, for instance, a certain medicine causes fewer hospital visits.
Psychology: Psychologists consider how interventions (like new therapy techniques) affect behavior. They also pay close attention to individual differences—why a method might help some people and not others.
Statistics: Statisticians develop tools (like matching or “controlling for” certain variables) to compare similar groups that differ mainly in whether they got the treatment or not.
Design and Analysis

We often use special designs (like randomized experiments, if possible) or, when experiments aren’t possible, statistical adjustments (like matching methods) to create fair comparisons.
The aim is to rule out confounding factors—other variables that might explain the difference between “treated” and “untreated” groups.
Causal Discovery (Beyond Just One Treatment)

Sometimes we don’t just want to measure the effect of one particular change; instead, we want to figure out how different factors interact in a whole system. Methods like path analysis and graphical models try to map out the relationships among many variables.
Think of these techniques as giving us a diagram of what causes what—like a flowchart of how different factors lead to, say, better test scores or improved health.
A Unifying Idea

Underlying all of this is the idea of counterfactual reasoning: asking what would have happened under different scenarios. Whether you call it “potential outcomes” (common in statistics), “structural equations” (common in economics), or “graphical models” (common in computer science), the goal is the same: to formalize the idea of cause-and-effect so we can measure it carefully and check if our assumptions make sense.
In short, counterfactual causal inference is about using data—and sometimes carefully designed studies—to figure out what really causes what, by imagining different “what if?” scenarios. Different fields may use different words or tools, but they share the core goal of moving beyond simple correlations and toward understanding genuine cause-and-effect relationships.







You said:
if I wanted to use R to model something like this, how could I approach it? I want to benchmark it through the command line, what are some tips for structuring it in a file instead of in RStudio

Reasoned for 5 seconds
Below is a high-level guide on how you might structure and run a causal inference project in R from the command line (rather than interactively in RStudio). This includes some practical tips and a minimal example for structuring your script.

1. Choosing a Causal Inference Approach
Depending on your study design and data availability, you might choose different R packages or methods:

Matching/Propensity Scores: Packages like MatchIt, twang, or Matching help in matching treated and control units to reduce confounding.
Instrumental Variables (IV): Packages like AER or ivreg for instrumental variable approaches (common in economics).
G-Computation / Inverse Probability Weighting (IPW): Packages like ipw or WeightIt.
Directed Acyclic Graphs (DAGs): Packages like dagitty for specifying causal structures.
Bayesian Approaches: Packages like rstan or brms, where you can incorporate prior knowledge.
For a simple demonstration, let’s assume we’re doing a propensity score matching analysis using the MatchIt package.

2. Structuring Your R Script
Header / Setup

Load libraries (library(MatchIt), etc.).
Load or parse command-line arguments if needed.
Set any global options (e.g., options(stringsAsFactors = FALSE)).
Data Loading

Read in your dataset (from CSV, database, or other file).
Perform any needed data cleaning or transformations.
Define Functions (Optional)

If your analysis has multiple steps, break it into functions.
For benchmarking or performance testing, you might wrap each method in a function you can time with system.time() or the microbenchmark package.
Model Fitting / Analysis

Specify your causal model.
Run the matching or weighting procedure.
Fit outcome models if necessary.
Results / Diagnostics

Print or store summary statistics.
Do any balance checking or diagnostic plots.
Export results (e.g., to CSV or to an RData file).
Wrap Up

Possibly clean up or remove large objects.
Save logs or final model objects as needed.
Example: my_causal_study.R
r
Copy
Edit
#!/usr/bin/env Rscript

# 1. Header / Setup
library(MatchIt)      # For matching
library(optparse)     # For command-line argument parsing (optional)

# If you want to parse command-line arguments, you can do something like:
option_list <- list(
  make_option(c("-i", "--input"), type="character", default="data.csv",
              help="Path to input CSV file", metavar="character")
)
opt_parser  <- OptionParser(option_list=option_list)
opt         <- parse_args(opt_parser)

# 2. Data Loading
# Example reading from CSV; uses --input argument if provided
data <- read.csv(opt$input)

# 3. (Optional) Functions for analysis or repeated tasks
balance_summary <- function(matchit_object) {
  summary(matchit_object)$sum.matched
}

# 4. Model Fitting / Analysis
# Suppose we have a binary treatment variable "treatment"
# and a set of covariates X1, X2, X3, etc.

match_formula <- as.formula("treatment ~ X1 + X2 + X3")

# Perform nearest-neighbor matching on the propensity score
m.out <- matchit(match_formula, data=data, method="nearest", distance="logit")

# 5. Results / Diagnostics
matched_data <- match.data(m.out)

# Print summary balance
cat("Summary of matching balance:\n")
print(balance_summary(m.out))

# Fit a simple outcome model (just as an example)
# Suppose "outcome" is our dependent variable
outcome_model <- lm(outcome ~ treatment + X1 + X2 + X3, data=matched_data)
cat("\nOutcome model summary (matched data):\n")
print(summary(outcome_model))

# 6. Wrap Up
# Save results to file if desired
saveRDS(m.out, file="matchit_object.rds")
saveRDS(outcome_model, file="outcome_model.rds")

cat("\nAnalysis complete.\n")
Note:

The script is written so it can be called from the command line (using Rscript my_causal_study.R --input mydata.csv).
#!/usr/bin/env Rscript on the first line (the “shebang”) allows Unix-like systems to execute this file directly if it’s marked as executable. On Windows, you can omit that line or just call it via Rscript.
3. Running & Benchmarking from the Command Line
Make it Executable (Linux/Mac)

bash
Copy
Edit
chmod +x my_causal_study.R
./my_causal_study.R --input your_data.csv
Standard Way (All Systems)

bash
Copy
Edit
Rscript my_causal_study.R --input your_data.csv
Benchmarking

If you just want rough timing, you can prefix your command with the shell’s time command in Linux/Mac:
bash
Copy
Edit
time Rscript my_causal_study.R --input your_data.csv
For more detailed profiling (e.g., to see where your script spends the most time), you can use R’s built-in Rprof() or the profvis package. You’d modify your script to wrap the parts you want to profile, or run in an interactive environment as needed.
4. Tips for Structuring & Good Practices
Keep Code Modular

Break your script into functions if it grows large. This makes it easier to read, test, and reuse parts of your code.
Use Logging

Include cat() or message() calls to track progress in the console, especially during long-running tasks.
Version Control

Maintain your R script in a version control system (e.g., Git). This is especially helpful for collaborative or long-term projects.
Command-Line Arguments

Use optparse, argparse, or docopt packages to handle command-line parameters (e.g., input filenames, hyperparameters for models, etc.). This makes your script more flexible.
Reproducibility

Set a seed for random processes (e.g., set.seed(123)) so that your results are reproducible.
Document your session info (e.g., using sessionInfo()) or store it in logs for debugging later.
Separation of Concerns

Keep data wrangling separate from modeling in your scripts if possible (two scripts or two sections). This helps you pinpoint where errors occur and also helps if you reuse the data wrangling for other analyses.
Summary
By laying out your script with clear sections for loading libraries, reading data, modeling, and exporting results, you can seamlessly run and benchmark causal inference analyses from the command line. You’ll find that the same workflow you might do interactively in RStudio simply requires more explicit logging and a bit of command-line argument handling so that it can run unattended and reproducibly.







You said:
how can I build a shell script to test different X values in this case?

Reasoned about scripting for testing values for 4 seconds
Below is a simple illustration of how you can create a shell script (bash example) to run your R script multiple times, each time testing a different set of covariates (i.e., different “X” variables). The idea is:

You define an array (or list) of possible sets of covariates in the shell script.
You loop through those sets and call your R script each time, passing the chosen set of covariates as a command-line argument.
Your R script then picks up that argument, plugs it into the model, and runs the analysis.
Step 1: Modify Your R Script to Accept Covariates
Let’s adapt the example my_causal_study.R so it can read a --covariates argument that specifies which X variables to use:

r
Copy
Edit
#!/usr/bin/env Rscript

library(MatchIt)
library(optparse)

# Define command-line options
option_list <- list(
  make_option(c("-i", "--input"), type="character", default="data.csv",
              help="Path to input CSV file", metavar="character"),
  make_option(c("-c", "--covariates"), type="character", default="X1+X2+X3",
              help="Covariates for the matching formula, e.g. 'X1+X2+X3'", metavar="character")
)

opt_parser  <- OptionParser(option_list=option_list)
opt         <- parse_args(opt_parser)

# Read the data
data <- read.csv(opt$input)

# Build the formula string; 'treatment' is the outcome to be predicted
# Note we paste "treatment ~" with the user's input
match_formula_str <- paste("treatment ~", opt$covariates)
match_formula     <- as.formula(match_formula_str)

# Do the matching
m.out <- matchit(match_formula, data=data, method="nearest", distance="logit")

# Grab matched data
matched_data <- match.data(m.out)

# Fit an outcome model, e.g. 'outcome ~ treatment + <covariates>'
outcome_formula_str <- paste("outcome ~ treatment +", opt$covariates)
outcome_formula     <- as.formula(outcome_formula_str)
outcome_model       <- lm(outcome_formula, data=matched_data)

# Print some results
cat("=== Matching Formula ===\n", match_formula_str, "\n\n")
cat("=== Balance Summary ===\n")
print(summary(m.out)$sum.matched)

cat("\n=== Outcome Model Summary ===\n")
print(summary(outcome_model))
Explanation
We added a new command-line argument --covariates (short version -c).
We construct the matching formula dynamically based on whatever is passed in.
You can then call Rscript my_causal_study.R --covariates "X1+X2+X3" to specify which variables to use.
Step 2: Create a Shell Script to Loop Over Different Covariates
Below is a minimal bash script, let’s call it test_covariates.sh. It defines a list of different covariate combinations you want to test, then runs the R script for each one.

bash
Copy
Edit
#!/bin/bash

# Make sure your script is executable:
# chmod +x test_covariates.sh

# Array (or list) of different covariate combinations you want to try.
covariate_list=(
  "X1+X2+X3"
  "X1+X2+X3+X4"
  "X2+X3+X4+X5"
  "X1+X4"
)

# Path to your R script
R_SCRIPT="my_causal_study.R"

# Path to your data
DATA_FILE="data.csv"

# Loop over each covariate combination
for covs in "${covariate_list[@]}"; do
  echo "============================"
  echo "Testing covariates: $covs"
  echo "============================"

  # Optionally time the run (on Mac/Linux):
  time Rscript "$R_SCRIPT" \
    --input "$DATA_FILE" \
    --covariates "$covs"

  # You could also redirect output to a log file if you want:
  # Rscript "$R_SCRIPT" --input "$DATA_FILE" --covariates "$covs" > "log_${covs}.txt" 2>&1
done
Explanation
We store different formulas (strings) in an array (covariate_list).
For each set of covariates, we call the R script, passing --covariates "$covs".
We also pass the data file with --input "$DATA_FILE".
We use the shell built-in command time for rough benchmarking.
You can capture console output to a log file if you like by redirecting output (>) and errors (2>&1).
Step 3: Run the Shell Script
Make it executable (once):
bash
Copy
Edit
chmod +x test_covariates.sh
Run it:
bash
Copy
Edit
./test_covariates.sh
This will sequentially run your R script with each covariate combination.
Step 4: Tips & Extensions
Storing Outputs

You might store each run’s results (like matched data or model summaries) in separate files. For example, you can pass an additional --output parameter to your R script and change file names based on $covs.
Profiling

Use time or more advanced profiling (Rprof or profvis) if you want to see which part of the code is slow.
Parallelization

If your dataset is large or you have many covariate sets, consider running each in parallel (using GNU Parallel or a job scheduler).
Logging

For clarity, you might name your log files with a prefix that includes the covariate string. For instance,
bash
Copy
Edit
Rscript "$R_SCRIPT" --input "$DATA_FILE" --covariates "$covs" \
  > "logs/out_${covs}.log" 2>&1
Multiple Data Sources

If you have multiple data files, you can also loop through data files in your shell script. You’d have an outer loop for data files, and an inner loop for covariates.
Putting It All Together
R Script: my_causal_study.R (handles --covariates).
Shell Script: test_covariates.sh (loops over covariates, calls the R script).
That’s it! You can now systematically test different sets of “X” variables (covariates) from the command line, capturing timing information and output logs along the way.
```
