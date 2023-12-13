// chartparser/chartparser.go

package chartparser

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

// Chart represents the structure of Chart.yaml
type Chart struct {
	ApiVersion string `yaml:"apiVersion"`
	Name       string `yaml:"name"`
	Version    string `yaml:"version"`
}

// Dictionary holds a map of chart names to their versions
type Dictionary struct {
	Charts map[string]string
}

// NewDictionary creates a new Dictionary
func NewDictionary() *Dictionary {
	return &Dictionary{Charts: make(map[string]string)}
}

// PopulateFromWalk traverses the directory tree and fills the chart dictionary
func (d *Dictionary) PopulateFromWalk(rootPath string) error {
	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		baseName := strings.ToLower(filepath.Base(path))
		//baseName := strings.ToLower(filepath.Base(path))
		switch baseName {
		case "chart.yaml", "chart.yml":
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			var chart Chart
			if err = yaml.Unmarshal(data, &chart); err != nil {
				return err
			}
			d.Charts[chart.Name] = chart.Version
		}
		return nil
	})
}
