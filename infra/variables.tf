variable "name" {
  type        = string
  description = "the name of the application to be deployed"
}

variable "location" {
  type        = string
  description = "The location where resources will be created"
}

variable "network" {
  type = object({
    network_address_space = list(string)
    db_subnet             = list(string)
    jb_subnet             = list(string)
  })
}

variable "azurerm_postgresql_server" {
  type = object({
    name                          = string
    administrator_login           = string
    sku_name                      = string
    version                       = string
    storage_mb                    = number
    backup_retention_days         = number
    geo_redundant_backup_enabled  = bool
    auto_grow_enabled             = bool
    public_network_access_enabled = bool
  })
  description = "https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_flexible_server"
}

variable "jumpboxvm" {
  type = object({
    vm_size                     = string
    storage_image_reference_sku = string
    os_profile_admin_username   = string
    ssh_keys_path               = string
  })
  description = "https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_machine"
}