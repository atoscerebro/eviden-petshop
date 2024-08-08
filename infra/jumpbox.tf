
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

resource "azurerm_virtual_machine" "jumpbox_vm" {
  name                  = "${var.name}-jumpboxvm"
  location              = var.location
  resource_group_name   = azurerm_resource_group.petshop_rg.name
  network_interface_ids = [azurerm_network_interface.jumpbox_nic.id]
  vm_size               = var.jumpboxvm.vm_size

  storage_os_disk {
    name              = "${var.name}-osdisk"
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  storage_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = var.jumpboxvm.storage_image_reference_sku
    version   = "latest"
  }

  os_profile {
    computer_name  = "${var.name}-jumpboxvm"
    admin_username = var.jumpboxvm.os_profile_admin_username
  }

  os_profile_linux_config {
    disable_password_authentication = true
    ssh_keys {
      path     = "/home/${var.jumpboxvm.os_profile_admin_username}/.ssh/authorized_keys"
      key_data = file(var.jumpboxvm.ssh_keys_path)
    }
  }
}

resource "azurerm_subnet_network_security_group_association" "jumpbox_netsec_link" {
  subnet_id                 = azurerm_subnet.jb_subnet.id
  network_security_group_id = azurerm_network_security_group.main.id
}