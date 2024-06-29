terraform {
  required_version = ">= 1.1.0"

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0.2"
    }
    
    azuread = {
      source  = "hashicorp/azuread"
      version = "2.41.0"
    }
  }
  
  cloud {
    organization = "the-junkyard"

    workspaces {
      name = "ehealth-microservices-hcp"
    }
  }
}