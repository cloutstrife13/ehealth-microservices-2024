# eHealth Microservices

> **DISCLAIMER:** This project is built purely for educational purposes to learn Go (Golang), Kubernetes microservices, and Bazel.

This project simulates an eHealth platform designed to automate and streamline digital workflows aligned with the German (Bismarck) healthcare system. The goal is to eliminate paper-based bureaucracy by allowing patients to securely register, verify their identity out-of-band, book appointments, order prescriptions, and manage their health data.

---

## Prerequisites

Ensure you have the required CLI tools installed on your machine:

```bash
# Install Go (Go 1.24+ recommended)
brew install go

# Install Docker Engine
brew install --cask docker

# Install Bazel build system
brew install bazel

# Install Kubernetes & Local Cluster tooling
brew install kubectl minikube skaffold
```

## First-Time Setup & Troubleshooting
If you are cloning this repository for the first time, ensure your local environment is clean to prevent toolchain or parent workspace conflicts:

```bash
# Disable Go workspace mode if a parent go.work file exists
export GOWORK=off

# Clear legacy Bazel build caches before your initial build
bazel clean --expunge
```

## Dependency Management
This repository uses standard Go modules integrated with Gazelle to manage Bazel BUILD.bazel target definitions automatically.

```bash
# Tidy up Go modules and track transitive requirements
go mod tidy

# Add a new dependency via Go CLI
go get <dependency-name>

# Auto-generate and update Bazel BUILD.bazel files across the workspace
bazel run //:gazelle

# Sync and update Bazel BUILD targets across the workspace
bazel run //:gazelle-update
```

## Launching Microservices
Start your local Minikube cluster and deploy the microservice architecture with hot-reloading enabled via Skaffold:

```bash
# Start Minikube cluster and enable the Ingress addon
minikube start
minikube addons enable ingress

# Deploy microservices to K8s and watch logs in real time
skaffold dev
```

## Running Tests
Execute test suites using Bazel:

```bash
# Run all unit and integration tests across the workspace
bazel test --test_output=errors //...

# Run specific service tests (e.g., patient-service)
bazel test --test_output=errors //services/patient-service/test/patients/...

# Run tests with verbose, line-by-line Ginkgo output
bazel test --test_output=all --test_arg=--ginkgo.v //services/patient-service/test/patients/...
```
