resource "singlestore_billing" "my_billing" {
  aggregate_by = "month"
  end_time     = "...my_end_time..."
  metric       = "...my_metric..."
  start_time   = "...my_start_time..."
}