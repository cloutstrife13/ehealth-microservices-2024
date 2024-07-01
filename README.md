# ehealth-microservice-2024

DISCLAIMER: This project is just for fun to learn Golang.

This is an eHealth project geared towards automating the Bismark system in Germany such that the need of paper-based bureaucracy is eliminated, allowing patients to easily book appointments, order prescriptions, and manage their data.

## Prerequisite

```bash
# Install golang
brew install go

# Install Docker
brew cask install docker

# Install bazel
brew install bazel

# Install kubectl
brew install kubectl

# Install minikube
brew install minikube

# Install skaffold
brew install skaffold
```

## Dependency management

```bash
# Add some dependency with Go CLI
go get <dependency-name>

# Update dependency list in Bazel
bazel run //:gazelle-update
```

## Launching microservices

```bash
skaffold dev
```
