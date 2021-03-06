<!--- GENERATED BY gomplate from scripts/docs/monitor-page.md.tmpl --->

# collectd/disk

This monitor collects information about the usage of
physical disks and logical disks (partitions).

See https://collectd.org/wiki/index.php/Plugin:Disk.


Monitor Type: `collectd/disk`

[Monitor Source Code](https://github.com/signalfx/signalfx-agent/tree/master/internal/monitors/collectd/disk)

**Accepts Endpoints**: No

**Multiple Instances Allowed**: **No**

## Configuration

| Config option | Required | Type | Description |
| --- | --- | --- | --- |
| `disks` | no | `list of strings` | Which devices to include/exclude (**default:** `[/^loop[0-9]+$/ /^dm-[0-9]+$/]`) |
| `ignoreSelected` | no | `bool` | If true, the disks selected by `disks` will be excluded and all others included. (**default:** `true`) |




## Metrics

The following table lists the metrics available for this monitor. Metrics that are marked as Included are standard metrics and are monitored by default.

| Name | Type | Included | Description |
| ---  | ---  | ---    | ---         |
| `disk_io_time.io_time` | cumulative |  | Amount of time spent doing IO in ms |
| `disk_io_time.weighted_io_time` | cumulative |  | Amount of time spent doing IO in ms multiplied by the queue length |
| `disk_merged.read` | cumulative |  | The number of disk reads merged into single physical disk access operations. |
| `disk_merged.write` | cumulative |  | The number of disk writes merged into single physical disk access operations. |
| `disk_octets.read` | cumulative |  | The number of bytes (octets) read from a disk. |
| `disk_octets.write` | cumulative |  | The number of bytes (octets) written to a disk. |
| `disk_ops.read` | cumulative | ✔ | The number of disk read operations. |
| `disk_ops.write` | cumulative | ✔ | The number of disk write operations. |
| `disk_time.read` | cumulative |  | The average amount of time it took to do a read operation. |
| `disk_time.write` | cumulative |  | The average amount of time it took to do a write operation. |
| `pending_operations` | gauge |  | Number of pending operations |


To specify custom metrics you want to monitor, add a `metricsToInclude` filter
to the agent configuration, as shown in the code snippet below. The snippet
lists all available custom metrics. You can copy and paste the snippet into
your configuration file, then delete any custom metrics that you do not want
sent.

Note that some of the custom metrics require you to set a flag as well as add
them to the list. Check the monitor configuration file to see if a flag is
required for gathering additional metrics.

```yaml

metricsToInclude:
  - metricNames:
    - disk_io_time.io_time
    - disk_io_time.weighted_io_time
    - disk_merged.read
    - disk_merged.write
    - disk_octets.read
    - disk_octets.write
    - disk_time.read
    - disk_time.write
    - pending_operations
    monitorType: collectd/disk
```




