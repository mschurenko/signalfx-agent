<!--- GENERATED BY gomplate from scripts/docs/monitor-page.md.tmpl --->

# collectd/processes

Gathers information about processes running on
the host.  See
https://collectd.org/documentation/manpages/collectd.conf.5.shtml#plugin_processes
and https://collectd.org/wiki/index.php/Plugin:Processes for more
information on the configuration options.

Example:

```yaml
 procPath: /proc
 monitors:
  - type: collectd/processes
    processes:
      - mysql
      - myapp
    processMatch:
      docker: "docker.*"
    collectContextSwitch: true
```

The above config will send process metrics for processes named *mysql* and
*myapp*, along with additional metrics on the number of context switches the
process has made.  Also, all processes that start with `docker` will have
their process metrics aggregated together and sent with a `plugin_instance`
value of `docker`.


Monitor Type: `collectd/processes`

[Monitor Source Code](https://github.com/signalfx/signalfx-agent/tree/master/internal/monitors/collectd/processes)

**Accepts Endpoints**: No

**Multiple Instances Allowed**: **No**

## Configuration

| Config option | Required | Type | Description |
| --- | --- | --- | --- |
| `processes` | no | `list of strings` | A list of process names to match |
| `processMatch` | no | `map of strings` | A map with keys specifying the `plugin_instance` value to be sent for the values which are regexes that match process names.  See example in description. |
| `collectContextSwitch` | no | `bool` | Collect metrics on the number of context switches made by the process (**default:** `false`) |
| `procFSPath` | no | `string` | (Deprecated) Please set the agent configuration `procPath` instead of this monitor configuration option. The path to the proc filesystem -- useful to override if the agent is running in a container. |




## Metrics

The following table lists the metrics available for this monitor. Metrics that are marked as Included are standard metrics and are monitored by default.

| Name | Type | Included | Description |
| ---  | ---  | ---    | ---         |
| `disk_octets.read` | cumulative |  |  |
| `disk_octets.write` | cumulative |  |  |
| `fork_rate` | cumulative |  |  |
| `io_octets.rx` | cumulative |  |  |
| `io_octets.tx` | cumulative |  |  |
| `io_ops.read` | cumulative |  |  |
| `io_ops.write` | cumulative |  |  |
| `ps_code` | gauge |  |  |
| `ps_count.processes` | gauge |  |  |
| `ps_count.threads` | gauge |  |  |
| `ps_cputime.syst` | cumulative |  |  |
| `ps_cputime.user` | cumulative |  |  |
| `ps_data` | gauge |  |  |
| `ps_pagefaults.majflt` | cumulative |  |  |
| `ps_pagefaults.minflt` | cumulative |  |  |
| `ps_rss` | gauge |  |  |
| `ps_stacksize` | gauge |  |  |
| `ps_state.blocked` | gauge |  |  |
| `ps_state.paging` | gauge |  |  |
| `ps_state.running` | gauge |  |  |
| `ps_state.sleeping` | gauge |  |  |
| `ps_state.stopped` | gauge |  |  |
| `ps_state.zombies` | gauge |  |  |
| `ps_vm` | gauge |  |  |





