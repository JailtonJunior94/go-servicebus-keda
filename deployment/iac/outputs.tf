output "connection_string_sb" {
  value     = nonsensitive(azurerm_servicebus_namespace.keda_poc_servicebus_namespace.default_primary_connection_string)
}