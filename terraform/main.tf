provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name
  location = "westeurope"

  tags = {
    Environment = "ehealth-microservices-dev"
    Team = "infrastructure"
  }
}

# Create a virtual network
resource "azurerm_virtual_network" "vnet" {
  name                = "ehealth-microservices-vpc"
  address_space       = ["10.0.0.0/16"]
  location            = "westeurope"
  resource_group_name = azurerm_resource_group.rg.name
}
