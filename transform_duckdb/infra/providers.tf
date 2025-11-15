terraform {
  required_version = "= 1.10.5"

    required_providers {
      azurerm = {
        source  = "hashicorp/azurerm"
        version = "4.46.0"
      }
    }
}

provider "azurerm" {
    features {}
    subscription_id = "4b72a73f-e970-40e1-b041-499eebd327a7"
    tenant_id       = "55226c2c-0b83-4621-a5cd-e8e0e57ec920"
}