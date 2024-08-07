variable "name" {
  type        = string
  description = "the name of the application to be deployed"
}

variable "location" {
  type        = string
  description = "The location where resources will be created"
}

variable "azurerm_postgresql_server" {
  type = object({
    name                             = string
    administrator_login              = string
    sku_name                         = string
    version                          = string
    storage_mb                       = number
    backup_retention_days            = number
    geo_redundant_backup_enabled     = bool
    auto_grow_enabled                = bool
    public_network_access_enabled    = bool
    ssl_enforcement_enabled          = bool
    ssl_minimal_tls_version_enforced = string
  })
  description = "https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_server"
}