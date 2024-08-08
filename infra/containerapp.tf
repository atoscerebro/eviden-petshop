
# resource "azurerm_container_app_environment" "petshop_app_environment" {
#   name                = "${var.name}-env"
#   resource_group_name = azurerm_resource_group.petshop_rg.name
#   location            = var.location
# }

# resource "azurerm_container_app" "petshop_application" {
#   name                = var.name
#   resource_group_name = azurerm_resource_group.petshop_rg.name

#   revision_mode                = "Single"
#   container_app_environment_id = azurerm_container_app_environment.petshop_app_environment.id

#   template {
#     container {
#       name   = "petshop"
#       image  = "<TODO>"
#       cpu    = "2"
#       memory = "4Gi"
#     }
#   }
# }
