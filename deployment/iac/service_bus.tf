resource "azurerm_servicebus_namespace" "keda_poc_servicebus_namespace" {
  name                = "keda-poc-ns"
  location            = azurerm_resource_group.poc_keda_rg.location
  resource_group_name = azurerm_resource_group.poc_keda_rg.name
  sku                 = "Basic"
}

resource "azurerm_servicebus_queue" "keda_poc_queue" {
  name         = "keda-poc-queue"
  namespace_id = azurerm_servicebus_namespace.keda_poc_servicebus_namespace.id
}
