resource "azurerm_resource_group" "petshop_rg" {
  name     = "${var.name}-rg"
  location = var.location
}