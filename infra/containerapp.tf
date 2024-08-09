# resource "azurerm_container_registry" "container_registry" {
#   name                = "examplecontainerregistry"
#   resource_group_name = azurerm_resource_group.petshop_rg.name
#   location            = var.location
#   sku                 = "Basic"

#   admin_enabled = true
# }

# resource "azurerm_container_app_environment" "petshop_containerapp_environment" {
#   name                = "${var.name}-app-env"
#   resource_group_name = azurerm_resource_group.petshop_rg.name
#   location            = var.location

#   tags = {
#     environment = var.petshop_containerapp_environment.tags
#   }
# }

# resource "azurerm_container_app" "petshop_containerapp" {
#   name                = "${var.name}-aca"
#   resource_group_name =  azurerm_resource_group.petshop_rg.name
#   location            =  var.location
#   container_app_environment_id = azurerm_container_app_environment.petshop_containerapp_environment.id

#   revision_mode = var.petshop_containerapp.revision_mode

#   template {
#     container {
#       name   = "${var.name}-container"
#       image  = "${azurerm_container_registry.container_registry.login_server}/${var.name}:latest"
#       cpu    = var.petshop_containerapp.container.cpu
#       memory = var.petshop_containerapp.container.memory
#     }
#   }
#   ingress {
#     external_enabled = var.petshop_containerapp.ingress.external_enabled   
#     target_port      = var.petshop_containerapp.ingress.target_port
#     traffic_weight {
#       latest_revision = var.petshop_containerapp.ingress.traffic_weight.latest_revision
#       percentage      = var.petshop_containerapp.ingress.traffic_weight.percentage
#     }
#   }

#   tags = {
#     environment = var.petshop_containerapp.tags.environment
#   }
# }

# resource "azurerm_container_app_registry" "petshop_containerapp_registry" {
#   container_app_id = var.petshop_containerapp_registry.container_app_id
#   registry_id      = var.petshop_containerapp_registry.registry_id
#   username         = var.petshop_containerapp_registry.username
#   password         = var.petshop_containerapp_registry.password
# }
