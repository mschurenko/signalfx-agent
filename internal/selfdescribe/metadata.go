package selfdescribe

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const monitorMetadataFile = "metadata.yaml"

// MetricMetadata contains a metric's metadata.
type MetricMetadata struct {
	Alias       string  `json:"alias,omitempty"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Group       *string `json:"group"`
	Included    bool    `json:"included" default:"false"`
}

// PropMetadata contains a property's metadata.
type PropMetadata struct {
	Dimension   string `json:"dimension"`
	Description string `json:"description"`
}

// GroupMetadata contains a group's metadata.
type GroupMetadata struct {
	Description string `json:"description"`
}

// MonitorMetadata contains a monitor's metadata.
type MonitorMetadata struct {
	MonitorType string                    `json:"monitorType" yaml:"monitorType"`
	SendAll     bool                      `json:"sendAll" yaml:"sendAll"`
	Dimensions  map[string]DimMetadata    `json:"dimensions"`
	Doc         string                    `json:"doc"`
	Groups      map[string]GroupMetadata  `json:"groups"`
	Metrics     map[string]MetricMetadata `json:"metrics"`
	Properties  map[string]PropMetadata   `json:"properties"`
	// True if the list of metrics is definitively the set of metrics
	// this monitor will ever send. This impacts the sendExtraMetrics.
	MetricsExhaustive bool `json:"metricsExhaustive" yaml:"metricsExhaustive" default:"false"`
}

// PackageMetadata describes a package directory that may have one or more monitors.
type PackageMetadata struct {
	// Common is a section to allow multiple monitors to place shared data.
	Common map[string]interface{}
	// PackageDir is the directory to output the generated code if not the same directory as the metadata.yaml.
	PackageDir string `json:"packageDir" yaml:"packageDir"`
	Monitors   []MonitorMetadata
	// Name of the package in go. If not set defaults to the directory name.
	GoPackage *string `json:"goPackage" yaml:"goPackage"`
	// Filesystem path to the package directory.
	PackagePath string `json:"-" yaml:"-"`
	Path        string `json:"-" yaml:"-"`
}

// DimMetadata contains a dimension's metadata.
type DimMetadata struct {
	Description string `json:"description"`
}

// CollectMetadata loads metadata for all monitors located in root as well as any subdirectories.
func CollectMetadata(root string) ([]PackageMetadata, error) {
	var packages []PackageMetadata

	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || info.Name() != monitorMetadataFile {
			return nil
		}

		var pkg PackageMetadata

		if bytes, err := ioutil.ReadFile(path); err != nil {
			return errors.Wrapf(err, "unable to read metadata file %s", path)
		} else if err := yaml.UnmarshalStrict(bytes, &pkg); err != nil {
			return errors.Wrapf(err, "unable to unmarshal file %s", path)
		}

		pkg.PackagePath = filepath.Dir(path)
		pkg.Path = path

		packages = append(packages, pkg)

		return nil
	}); err != nil {
		return nil, err
	}

	return packages, nil
}
