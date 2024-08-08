resource "azurerm_postgresql_flexible_server" "petshop_db" {
  name                = "${var.name}-pgdb"
  location            = var.location
  resource_group_name = azurerm_resource_group.petshop_rg.name
  lifecycle {
    ignore_changes = [
      zone
    ]
  }

  administrator_login    = var.azurerm_postgresql_server.administrator_login
  administrator_password = random_password.password.result

  sku_name   = var.azurerm_postgresql_server.sku_name
  version    = var.azurerm_postgresql_server.version
  storage_mb = var.azurerm_postgresql_server.storage_mb

  backup_retention_days         = var.azurerm_postgresql_server.backup_retention_days
  geo_redundant_backup_enabled  = var.azurerm_postgresql_server.geo_redundant_backup_enabled
  auto_grow_enabled             = var.azurerm_postgresql_server.auto_grow_enabled
  delegated_subnet_id           = azurerm_subnet.db_subnet.id
  private_dns_zone_id           = azurerm_private_dns_zone.pethshop_db_dns.id
  public_network_access_enabled = var.azurerm_postgresql_server.public_network_access_enabled
}
