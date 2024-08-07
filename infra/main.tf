resource "azurerm_resource_group" "petshop_rg" {
  name     = "${var.name}-rg"
  location = var.location
}

resource "azurerm_postgresql_server" "eviden-petshop" {
  name                = "${var.name}-pgdb"
  location            = var.location
  resource_group_name = azurerm_resource_group.petshop_rg.name

  administrator_login          = var.azurerm_postgresql_server.administrator_login
  administrator_login_password = random_password.password.result

  sku_name   = var.azurerm_postgresql_server.sku_name
  version    = var.azurerm_postgresql_server.version
  storage_mb = var.azurerm_postgresql_server.storage_mb

  backup_retention_days        = var.azurerm_postgresql_server.backup_retention_days
  geo_redundant_backup_enabled = var.azurerm_postgresql_server.geo_redundant_backup_enabled
  auto_grow_enabled            = var.azurerm_postgresql_server.auto_grow_enabled

  public_network_access_enabled    = var.azurerm_postgresql_server.public_network_access_enabled
  ssl_enforcement_enabled          = var.azurerm_postgresql_server.ssl_enforcement_enabled
  ssl_minimal_tls_version_enforced = var.azurerm_postgresql_server.ssl_minimal_tls_version_enforced
}


resource "azurerm_container_app_environment" "petshop_app_environment" {
  name                = "${var.name}-env"
  resource_group_name = azurerm_resource_group.petshop_rg.name
  location            = var.location
}

resource "azurerm_container_app" "petshop_application" {
  name                = var.name
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
