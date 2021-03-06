package winservices

import (
	"context"

	"github.com/signalfx/signalfx-agent/internal/core/config"
	"github.com/signalfx/signalfx-agent/internal/monitors"
	"github.com/signalfx/signalfx-agent/internal/monitors/types"
)

func init() {
	monitors.Register(&monitorMetadata, func() interface{} { return &Monitor{} }, &Config{})
}

// Config for this monitor
type Config struct {
	config.MonitorConfig `acceptsEndpoints:"false"`
	// Names of services to monitor.  All services will be monitored if none are specified.
	ServiceNames []string `yaml:"serviceNames"`
}

// Monitor for Utilization
type Monitor struct {
	Output types.Output
	cancel context.CancelFunc
}

// Shutdown stops the metric sync
func (m *Monitor) Shutdown() {
	if m.cancel != nil {
		m.cancel()
	}
}
