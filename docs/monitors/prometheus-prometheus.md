<!--- GENERATED BY gomplate from scripts/docs/monitor-page.md.tmpl --->

# prometheus/prometheus

This monitor scrapes [Prmoetheus server's own internal
collector](https://prometheus.io/docs/prometheus/latest/getting_started/#configuring-prometheus-to-monitor-itself)
metrics from a Prometheus exporter and sends them to SignalFx.  It is a
wrapper around the [prometheus-exporter](./prometheus-exporter.md) monitor
that provides a restricted but expandable set of metrics.


Monitor Type: `prometheus/prometheus`

[Monitor Source Code](https://github.com/signalfx/signalfx-agent/tree/master/internal/monitors/prometheus/prometheus)

**Accepts Endpoints**: **Yes**

**Multiple Instances Allowed**: Yes

## Configuration

| Config option | Required | Type | Description |
| --- | --- | --- | --- |
| `host` | **yes** | `string` | Host of the exporter |
| `port` | **yes** | `integer` | Port of the exporter |
| `username` | no | `string` | Basic Auth username to use on each request, if any. |
| `password` | no | `string` | Basic Auth password to use on each request, if any. |
| `useHTTPS` | no | `bool` | If true, the agent will connect to the exporter using HTTPS instead of plain HTTP. (**default:** `false`) |
| `skipVerify` | no | `bool` | If useHTTPS is true and this option is also true, the exporter's TLS cert will not be verified. (**default:** `false`) |
| `metricPath` | no | `string` | Path to the metrics endpoint on the exporter server, usually `/metrics` (the default). (**default:** `/metrics`) |
| `sendAllMetrics` | no | `bool` | Send all the metrics that come out of the Prometheus exporter without any filtering.  This option has no effect when using the prometheus exporter monitor directly since there is no built-in filtering, only when embedding it in other monitors. (**default:** `false`) |




## Metrics

The following table lists the metrics available for this monitor. Metrics that are marked as Included are standard metrics and are monitored by default.

| Name | Type | Included | Description |
| ---  | ---  | ---    | ---         |
| `net_conntrack_dialer_conn_attempted_total` | cumulative |  | Total number of connections attempted by the given dialer a given name |
| `net_conntrack_dialer_conn_closed_total` | cumulative |  | Total number of connections closed which originated from the dialer of a given name |
| `net_conntrack_dialer_conn_established_total` | cumulative |  | Total number of connections successfully established by the given dialer a given name |
| `net_conntrack_dialer_conn_failed_total` | cumulative |  | Total number of connections failed to dial by the dialer a given name |
| `net_conntrack_listener_conn_accepted_total` | cumulative |  | Total number of connections opened to the listener of a given name |
| `net_conntrack_listener_conn_closed_total` | cumulative |  | Total number of connections closed that were made to the listener of a given name |
| `prometheus_api_remote_read_queries` | gauge |  | The current number of remote read queries being executed or waiting |
| `prometheus_build_info` | gauge |  | A metric with a constant '1' value labeled by version, revision, branch, and goversion from which prometheus was built |
| `prometheus_config_last_reload_success_timestamp_seconds` | gauge |  | Timestamp of the last successful configuration reload |
| `prometheus_config_last_reload_successful` | gauge |  | Whether the last configuration reload attempt was successful |
| `prometheus_engine_queries` | gauge |  | The current number of queries being executed or waiting |
| `prometheus_engine_queries_concurrent_max` | gauge |  | The max number of concurrent queries |
| `prometheus_engine_query_duration_seconds` | cumulative |  | Query timings |
| `prometheus_engine_query_duration_seconds_count` | cumulative |  | Query timings (count) |
| `prometheus_http_request_duration_seconds` | cumulative |  | Histogram of latencies for HTTP requests |
| `prometheus_http_request_duration_seconds_bucket` | cumulative |  | Histogram of latencies for HTTP requests in the respective bucket |
| `prometheus_http_request_duration_seconds_count` | cumulative |  | Histogram of latencies for HTTP requests (count) |
| `prometheus_http_response_size_bytes` | cumulative |  | Histogram of response size for HTTP requests |
| `prometheus_http_response_size_bytes_bucket` | cumulative |  | Histogram of response size for HTTP requests in the respective bucket |
| `prometheus_http_response_size_bytes_count` | cumulative |  | Histogram of response size for HTTP requests |
| `prometheus_notifications_alertmanagers_discovered` | gauge |  | The number of alertmanagers discovered and active |
| `prometheus_notifications_dropped_total` | cumulative |  | Total number of alerts dropped due to errors when sending to Alertmanager |
| `prometheus_notifications_queue_capacity` | gauge |  | The capacity of the alert notifications queue |
| `prometheus_notifications_queue_length` | gauge |  | The number of alert notifications in the queue |
| `prometheus_rule_evaluation_duration_seconds` | cumulative |  | The duration for a rule to execute |
| `prometheus_rule_evaluation_duration_seconds_count` | cumulative |  | The duration for a rule to execute (count) |
| `prometheus_rule_evaluation_failures_total` | cumulative |  | The total number of rule evaluation failures |
| `prometheus_rule_group_duration_seconds` | cumulative |  | The duration of rule group evaluations |
| `prometheus_rule_group_duration_seconds_count` | cumulative |  | The duration of rule group evaluations (count) |
| `prometheus_rule_group_interval_seconds` | gauge |  | The interval of a rule group |
| `prometheus_rule_group_iterations_missed_total` | cumulative |  | The total number of rule group evaluations missed due to slow rule group evaluation |
| `prometheus_rule_group_iterations_total` | cumulative |  | The total number of scheduled rule group evaluations, whether executed or missed |
| `prometheus_rule_group_last_duration_seconds` | gauge |  | The duration of the last rule group evaluation |
| `prometheus_sd_azure_refresh_duration_seconds` | cumulative |  | The duration of a Azure-SD refresh in seconds |
| `prometheus_sd_azure_refresh_duration_seconds_count` | cumulative |  | The duration of a Azure-SD refresh in seconds (count) |
| `prometheus_sd_azure_refresh_failures_total` | cumulative |  | Number of Azure-SD refresh failures |
| `prometheus_sd_configs_failed_total` | cumulative |  | Total number of service discovery configurations that failed to load |
| `prometheus_sd_consul_rpc_duration_seconds` | cumulative |  | The duration of a Consul RPC call in seconds |
| `prometheus_sd_consul_rpc_duration_seconds_count` | cumulative |  | The duration of a Consul RPC call in seconds (count) |
| `prometheus_sd_consul_rpc_failures_total` | cumulative |  | The number of Consul RPC call failures |
| `prometheus_sd_discovered_targets` | gauge |  | Current number of discovered targets |
| `prometheus_sd_dns_lookup_failures_total` | cumulative |  | The number of DNS-SD lookup failures |
| `prometheus_sd_dns_lookups_total` | cumulative |  | The number of DNS-SD lookups |
| `prometheus_sd_ec2_refresh_duration_seconds` | cumulative |  | The duration of a EC2-SD refresh in seconds |
| `prometheus_sd_ec2_refresh_duration_seconds_count` | cumulative |  | The duration of a EC2-SD refresh in seconds (count) |
| `prometheus_sd_ec2_refresh_failures_total` | cumulative |  | The number of EC2-SD scrape failures |
| `prometheus_sd_file_read_errors_total` | cumulative |  | The number of File-SD read errors |
| `prometheus_sd_file_scan_duration_seconds` | cumulative |  | The duration of the File-SD scan in seconds |
| `prometheus_sd_file_scan_duration_seconds_count` | cumulative |  | The duration of the File-SD scan in seconds (count) |
| `prometheus_sd_gce_refresh_duration` | cumulative |  | The duration of a GCE-SD refresh in seconds |
| `prometheus_sd_gce_refresh_duration_count` | cumulative |  | The duration of a GCE-SD refresh in seconds (count) |
| `prometheus_sd_gce_refresh_failures_total` | cumulative |  | The number of GCE-SD refresh failures |
| `prometheus_sd_kubernetes_cache_last_resource_version` | gauge |  | Last resource version from the Kubernetes API |
| `prometheus_sd_kubernetes_cache_list_duration_seconds` | cumulative |  | Duration of a Kubernetes API call in seconds |
| `prometheus_sd_kubernetes_cache_list_duration_seconds_count` | cumulative |  | Duration of a Kubernetes API call in seconds (count) |
| `prometheus_sd_kubernetes_cache_list_items` | cumulative |  | Count of items in a list from the Kubernetes API |
| `prometheus_sd_kubernetes_cache_list_items_count` | cumulative |  | Count of items in a list from the Kubernetes API (count) |
| `prometheus_sd_kubernetes_cache_list_total` | cumulative |  | Total number of list operations |
| `prometheus_sd_kubernetes_cache_short_watches_total` | cumulative |  | Total number of short watch operations |
| `prometheus_sd_kubernetes_cache_watch_duration_seconds` | cumulative |  | Duration of watches on the Kubernetes API |
| `prometheus_sd_kubernetes_cache_watch_duration_seconds_count` | cumulative |  | Duration of watches on the Kubernetes API (count) |
| `prometheus_sd_kubernetes_cache_watch_events` | cumulative |  | Number of items in watches on the Kubernetes API |
| `prometheus_sd_kubernetes_cache_watch_events_count` | cumulative |  | Number of items in watches on the Kubernetes API (count) |
| `prometheus_sd_kubernetes_cache_watches_total` | cumulative |  | Total number of watch operations |
| `prometheus_sd_kubernetes_events_total` | cumulative |  | The number of Kubernetes events handled |
| `prometheus_sd_marathon_refresh_duration_seconds` | cumulative |  | The duration of a Marathon-SD refresh in seconds |
| `prometheus_sd_marathon_refresh_duration_seconds_count` | cumulative |  | The duration of a Marathon-SD refresh in seconds |
| `prometheus_sd_marathon_refresh_failures_total` | cumulative |  | The number of Marathon-SD refresh failures |
| `prometheus_sd_openstack_refresh_duration_seconds` | cumulative |  | The duration of an OpenStack-SD refresh in seconds |
| `prometheus_sd_openstack_refresh_duration_seconds_count` | cumulative |  | The duration of an OpenStack-SD refresh in seconds |
| `prometheus_sd_openstack_refresh_failures_total` | cumulative |  | The number of OpenStack-SD scrape failures |
| `prometheus_sd_received_updates_total` | cumulative |  | Total number of update events received from the SD providers |
| `prometheus_sd_triton_refresh_duration_seconds` | cumulative |  | The duration of a Triton-SD refresh in seconds |
| `prometheus_sd_triton_refresh_duration_seconds_count` | cumulative |  | The duration of a Triton-SD refresh in seconds |
| `prometheus_sd_triton_refresh_failures_total` | cumulative |  | The number of Triton-SD scrape failures |
| `prometheus_sd_updates_delayed_total` | cumulative |  | Total number of update events that couldn't be sent immediately |
| `prometheus_sd_updates_total` | cumulative |  | Total number of update events sent to the SD consumers |
| `prometheus_target_interval_length_seconds` | cumulative |  | Actual intervals between scrapes |
| `prometheus_target_interval_length_seconds_count` | cumulative |  | Actual intervals between scrapes |
| `prometheus_target_scrape_pool_sync_total` | cumulative |  | Total number of syncs that were executed on a scrape pool |
| `prometheus_target_scrapes_exceeded_sample_limit_total` | cumulative |  | Total number of scrapes that hit the sample limit and were rejected |
| `prometheus_target_scrapes_sample_duplicate_timestamp_total` | cumulative |  | Total number of samples rejected due to duplicate timestamps but different values |
| `prometheus_target_scrapes_sample_out_of_bounds_total` | cumulative |  | Total number of samples rejected due to timestamp falling outside of the time bounds |
| `prometheus_target_scrapes_sample_out_of_order_total` | cumulative |  | Total number of samples rejected due to not being out of the expected order |
| `prometheus_target_sync_length_seconds` | cumulative |  | Actual interval to sync the scrape pool |
| `prometheus_target_sync_length_seconds_count` | cumulative |  | Actual interval to sync the scrape pool |
| `prometheus_treecache_watcher_goroutines` | gauge |  | The current number of watcher goroutines |
| `prometheus_treecache_zookeeper_failures_total` | cumulative |  | The total number of ZooKeeper failures |
| `prometheus_tsdb_blocks_loaded` | gauge |  | Number of currently loaded data blocks |
| `prometheus_tsdb_checkpoint_creations_failed_total` | cumulative |  | Total number of checkpoint creations that failed |
| `prometheus_tsdb_checkpoint_creations_total` | cumulative |  | Total number of checkpoint creations attempted |
| `prometheus_tsdb_checkpoint_deletions_failed_total` | cumulative |  | Total number of checkpoint deletions that failed |
| `prometheus_tsdb_checkpoint_deletions_total` | cumulative |  | Total number of checkpoint deletions attempted |
| `prometheus_tsdb_compaction_chunk_range_seconds` | cumulative |  | Final time range of chunks on their first compaction |
| `prometheus_tsdb_compaction_chunk_range_seconds_bucket` | cumulative |  | Final time range of chunks on their first compaction in the respective bucket |
| `prometheus_tsdb_compaction_chunk_range_seconds_count` | cumulative |  | Final time range of chunks on their first compaction (count) |
| `prometheus_tsdb_compaction_chunk_samples` | cumulative |  | Final number of samples on their first compaction |
| `prometheus_tsdb_compaction_chunk_samples_bucket` | cumulative |  | Final number of samples on their first compaction in the respective bucket |
| `prometheus_tsdb_compaction_chunk_samples_count` | cumulative |  | Final number of samples on their first compaction (count) |
| `prometheus_tsdb_compaction_chunk_size_bytes` | cumulative |  | Final size of chunks on their first compaction |
| `prometheus_tsdb_compaction_chunk_size_bytes_bucket` | cumulative |  | Final size of chunks on their first compaction in the respective bucket |
| `prometheus_tsdb_compaction_chunk_size_bytes_count` | cumulative |  | Final size of chunks on their first compaction |
| `prometheus_tsdb_compaction_duration_seconds` | cumulative |  | Duration of compaction runs |
| `prometheus_tsdb_compaction_duration_seconds_bucket` | cumulative |  | Duration of compaction runs in the respective bucket |
| `prometheus_tsdb_compaction_duration_seconds_count` | cumulative |  | Duration of compaction runs (count) |
| `prometheus_tsdb_compactions_failed_total` | cumulative |  | Total number of compactions that failed for the partition |
| `prometheus_tsdb_compactions_total` | cumulative |  | Total number of compactions that were executed for the partition |
| `prometheus_tsdb_compactions_triggered_total` | cumulative |  | Total number of triggered compactions for the partition |
| `prometheus_tsdb_head_active_appenders` | gauge |  | Number of currently active appender transactions |
| `prometheus_tsdb_head_chunks` | gauge |  | Total number of chunks in the head block |
| `prometheus_tsdb_head_chunks_created_total` | cumulative |  | Total number of chunks created in the head |
| `prometheus_tsdb_head_chunks_removed_total` | cumulative |  | Total number of chunks removed in the head |
| `prometheus_tsdb_head_gc_duration_seconds` | cumulative |  | Runtime of garbage collection in the head block |
| `prometheus_tsdb_head_gc_duration_seconds_count` | cumulative |  | Runtime of garbage collection in the head block (count) |
| `prometheus_tsdb_head_max_time` | gauge |  | Maximum timestamp of the head block |
| `prometheus_tsdb_head_min_time` | gauge |  | Minimum time bound of the head block |
| `prometheus_tsdb_head_samples_appended_total` | cumulative |  | Total number of appended samples |
| `prometheus_tsdb_head_series` | gauge |  | Total number of series in the head block |
| `prometheus_tsdb_head_series_created_total` | cumulative |  | Total number of series created in the head |
| `prometheus_tsdb_head_series_not_found_total` | cumulative |  | Total number of requests for series that were not found |
| `prometheus_tsdb_head_series_removed_total` | cumulative |  | Total number of series removed in the head |
| `prometheus_tsdb_head_truncations_failed_total` | cumulative |  | Total number of head truncations that failed |
| `prometheus_tsdb_head_truncations_total` | cumulative |  | Total number of head truncations attempted |
| `prometheus_tsdb_lowest_timestamp` | gauge |  | Lowest timestamp value stored in the database |
| `prometheus_tsdb_reloads_failures_total` | cumulative |  | Number of times the database failed to reload block data from disk |
| `prometheus_tsdb_reloads_total` | cumulative |  | Number of times the database reloaded block data from disk |
| `prometheus_tsdb_retention_cutoffs_failures_total` | cumulative |  | Number of times the database failed to cut off block data from disk |
| `prometheus_tsdb_retention_cutoffs_total` | cumulative |  | Number of times the database cut off block data from disk |
| `prometheus_tsdb_symbol_table_size_bytes` | gauge |  | Size of symbol table on disk (in bytes) |
| `prometheus_tsdb_tombstone_cleanup_seconds` | cumulative |  | The time taken to recompact blocks to remove tombstones |
| `prometheus_tsdb_tombstone_cleanup_seconds_bucket` | cumulative |  | The time taken to recompact blocks to remove tombstones in the respective bucket |
| `prometheus_tsdb_tombstone_cleanup_seconds_count` | cumulative |  | The time taken to recompact blocks to remove tombstones (count) |
| `prometheus_tsdb_wal_completed_pages_total` | cumulative |  | Total number of completed pages |
| `prometheus_tsdb_wal_fsync_duration_seconds` | cumulative |  | Duration of WAL fsync |
| `prometheus_tsdb_wal_fsync_duration_seconds_count` | cumulative |  | Duration of WAL fsync (count) |
| `prometheus_tsdb_wal_page_flushes_total` | cumulative |  | Total number of page flushes |
| `prometheus_tsdb_wal_truncate_duration_seconds` | cumulative |  | Duration of WAL truncation |
| `prometheus_tsdb_wal_truncate_duration_seconds_count` | cumulative |  | Duration of WAL truncation (count) |
| `prometheus_tsdb_wal_truncations_failed_total` | cumulative |  | Total number of WAL truncations that failed |
| `prometheus_tsdb_wal_truncations_total` | cumulative |  | Total number of WAL truncations attempted |
| `promhttp_metric_handler_requests_in_flight` | gauge |  | Current number of scrapes being served |
| `promhttp_metric_handler_requests_total` | cumulative |  | Total number of scrapes by HTTP status code |





