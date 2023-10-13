resource "singlestore_private_connection" "my_privateconnection" {
  active_at           = "2023-09-18T09:56:56Z"
  allow_list          = "my-allow-list"
  created_at          = "2023-09-18T09:56:56Z"
  deleted_at          = "2023-09-18T09:56:56Z"
  outbound_allow_list = "arn:aws:iam:xxxxxxxxx:root"
  service_name        = "My service"
  status              = "ACTIVE"
  type                = "INBOUND"
  updated_at          = "2023-09-18T09:56:56Z"
  workspace_group_id  = "68af2f46-0000-1000-9000-3f6f5365d878"
  workspace_id        = "7d6b3a5a-0000-1000-9000-d454af4c785b"
}