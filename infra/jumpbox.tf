
resource "azurerm_network_interface" "jumpbox_nic" {
  name                = "${var.name}-nic"
  location            = var.location
  resource_group_name = azurerm_resource_group.petshop_rg.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.jb_subnet.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.main.id
  }
}

resource "azurerm_linux_virtual_machine" "jumpbox_vm" {
  name                = "${var.name}-jb"
  resource_group_name = azurerm_resource_group.petshop_rg.name
  location            = var.location
  size                = var.jumpboxvm.vm_size
  admin_username      = var.jumpboxvm.admin_username
  network_interface_ids = [
    azurerm_network_interface.jumpbox_nic.id,
  ]

  admin_ssh_key {
    username   = var.jumpboxvm.admin_username
    public_key = file(var.jumpboxvm.admin_ssh_key_public_key)
  }

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = var.jumpboxvm.source_image_reference_sku
    version   = "latest"
  }
}

resource "azurerm_subnet_network_security_group_association" "jumpbox_netsec_link" {
  subnet_id                 = azurerm_subnet.jb_subnet.id
  network_security_group_id = azurerm_network_security_group.main.id
}