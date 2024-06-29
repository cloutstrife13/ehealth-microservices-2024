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

resource "azurerm_virtual_network" "vnet" {
  name                = "ehealth-microservices-vpc"
  address_space       = ["10.0.0.0/16"]
  location            = "westeurope"
  resource_group_name = azurerm_resource_group.rg.name
}

resource "azurerm_aadb2c_directory" "tenant-rosaria" {
  country_code            = "IE"
  data_residency_location = "Europe"
  display_name            = "Rosaria"
  domain_name             = "rosarialtd.onmicrosoft.com"
  resource_group_name     = azurerm_resource_group.rg.name
  sku_name                = "PremiumP1"
}

resource "azurerm_aadb2c_directory" "tenant-sambreque" {
  country_code            = "IE"
  data_residency_location = "Europe"
  display_name            = "Sambreque"
  domain_name             = "sambrequeltd.onmicrosoft.com"
  resource_group_name     = azurerm_resource_group.rg.name
  sku_name                = "PremiumP1"
}