resource "azurerm_virtual_network" "main" {
  name                = "${var.name}-vnet"
  address_space       = var.network.network_address_space
  location            = var.location
  resource_group_name = azurerm_resource_group.petshop_rg.name
}

resource "azurerm_subnet" "db_subnet" {
  name                 = "${var.name}-db-subnet"
  resource_group_name  = azurerm_resource_group.petshop_rg.name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes     = var.network.db_subnet
  delegation {
    name = "${var.name}-fs"
    service_delegation {
      name = "Microsoft.DBforPostgreSQL/flexibleServers"
      actions = [
        "Microsoft.Network/virtualNetworks/subnets/join/action",
      ]
    }
  }
  depends_on = [azurerm_virtual_network.main]
}

resource "azurerm_subnet" "jb_subnet" {
  name                 = "${var.name}-jb-subnet"
  resource_group_name  = azurerm_resource_group.petshop_rg.name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes     = var.network.jb_subnet
  depends_on           = [azurerm_virtual_network.main]
}

resource "azurerm_network_security_group" "main" {
  name                = "${var.name}-nsg"
  location            = var.location
  resource_group_name = azurerm_resource_group.petshop_rg.name
}

resource "azurerm_network_security_rule" "ssh" {
  name                        = "Allow-SSH"
  priority                    = 1001
  direction                   = "Inbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_range      = "22"
  source_address_prefix       = "*"
  destination_address_prefix  = "*"
  resource_group_name         = azurerm_resource_group.petshop_rg.name
  network_security_group_name = azurerm_network_security_group.main.name
}

resource "azurerm_private_dns_zone" "pethshop_db_dns" {
  name                = "${var.name}.postgres.database.azure.com"
  resource_group_name = azurerm_resource_group.petshop_rg.name
}

resource "azurerm_private_dns_zone_virtual_network_link" "petshop_db_dns_link" {
  name                  = "${var.name}-dns-link"
  private_dns_zone_name = azurerm_private_dns_zone.pethshop_db_dns.name
  virtual_network_id    = azurerm_virtual_network.main.id
  resource_group_name   = azurerm_resource_group.petshop_rg.name
}

resource "azurerm_public_ip" "main" {
  name                = "${var.name}-ip"
  location            = var.location
  resource_group_name = azurerm_resource_group.petshop_rg.name
  allocation_method   = "Dynamic"
}