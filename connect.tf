resource "azurerm_virtual_hub_connection" "connect" {
  name                      = "eddf5c7f-8021-42a6-bf67-9a232373cd80"
  virtual_hub_id            = "/subscriptions/e0b91e0c-6967-4d2b-a458-bcfd7aff0871/resourceGroups/az-vwan-cloud2/providers/Microsoft.Network/virtualHubs/neu-hub01"
  remote_virtual_network_id = "/subscriptions/e0b91e0c-6967-4d2b-a458-bcfd7aff0871/resourceGroups/CostcoTesting/providers/Microsoft.Network/virtualNetworks/costcotestingvnet1"
}
