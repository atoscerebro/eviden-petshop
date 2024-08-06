resource "azurerm_resource_group" "petshop_rg" {
  name     = "eviden-petshop-rg"
  location = "UK South"
}

resource "azurerm_container_app_environment" "petshop_app_environment" {
  name                = "eviden-petshop-env"
  resource_group_name = azurerm_resource_group.petshop_rg.name
  location            = azurerm_resource_group.petshop_rg.location
}

resource "azurerm_container_app" "petshop_application" {
  name                = "eviden-petshop"
  resource_group_name = azurerm_resource_group.petshop_rg.name

  revision_mode                = "Single"
  container_app_environment_id = azurerm_container_app_environment.petshop_app_environment.id

  template {
    container {
      name   = "petshop"
      image  = "<TODO>"
      cpu    = "2"
      memory = "4Gi"
    }
  }
}
