# ehealth-microservice-2024

DISCLAIMER: This project is just for fun to learn Golang.

This is an eHealth project geared towards automating the Bismark system in Germany such that the need of paper-based bureaucracy is eliminated, allowing patients to easily book appointments, order prescriptions, and manage their data.

## Prerequisite

```bash
# Install golang (Go 1.24+ recommended)
brew install go

# Install Docker
brew install --cask docker

# Install bazel
brew install bazel

# Install kubectl
brew install kubectl

# Install minikube
brew install minikube

# Install skaffold
brew install skaffold
```

## First-Time Setup & Troubleshooting
If you are cloning this repository for the first time, ensure your environment is clean to prevent local toolchain or parent workspace conflicts:

```bash
# If you have a global go.work file in a parent directory that causes version conflicts:
export GOWORK=off

# Clear any legacy Bazel states before your initial build
bazel clean --expunge
```

## Dependency management

```bash
# Tidy up Go modules and track transitive requirements
go mod tidy

# Add some dependency with Go CLI
go get <dependency-name>

# Update dependency list in Bazel
bazel run //:gazelle-update
```

## Launching microservices

```bash
# Make sure your minikube cluster is running and ingress addon is enabled
minikube start
minikube addons enable ingress

# Deploy and start services in Kubernetes with hot-reload
skaffold dev
```

## Running tests

```bash
# Run all tests across the entire workspace
bazel test --test_output=errors //...

# Run specific integration tests (e.g., patient-service tests)
bazel test --test_output=errors //services/patient-service/test/patients/...

# Run specific integration tests with verbose, line-by-line test output
bazel test --test_output=all --test_arg=--ginkgo.v //services/patient-service/test/patients/...
```
