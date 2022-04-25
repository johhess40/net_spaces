resource "azurerm_virtual_hub_connection" "connect" {
  name                      = "e097cf9e-5552-4b1a-bfc8-a8d9eed7e36a"
  virtual_hub_id            = "/subscriptions/e0b91e0c-6967-4d2b-a458-bcfd7aff0871/resourceGroups/az-vwan-cloud2/providers/Microsoft.Network/virtualHubs/neu-hub01"
  remote_virtual_network_id = "/subscriptions/e0b91e0c-6967-4d2b-a458-bcfd7aff0871/resourceGroups/CostcoTesting/providers/Microsoft.Network/virtualNetworks/costcotestingvnet1"
}
output "connection" {
  value = azurerm_virtual_hub_connection.connect.id
}
