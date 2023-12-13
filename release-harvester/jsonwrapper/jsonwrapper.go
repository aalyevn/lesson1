package jsonwrapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// JSONDataWrapper holds the JSON data and provides methods to operate on it.
type JSONDataWrapper struct {
	Data map[string]interface{}
}

// New creates a new instance of JSONDataWrapper.
func New(filePath string) (*JSONDataWrapper, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}

	return &JSONDataWrapper{Data: jsonData}, nil
}

// UpdateChartVersionFromDictionary updates all chart versions in the JSON data using a dictionary of chart names to versions.
func (j *JSONDataWrapper) UpdateChartVersionFromDictionary(charts map[string]string, chartNamePath string, chartVersionPath string) error {
	for chartName, chartVersion := range charts {
		err := j.UpdateChartVersion(chartName, chartVersion, chartNamePath, chartVersionPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateChartVersion updates the chart version in the JSON data using provided paths
func (j *JSONDataWrapper) UpdateChartVersion(chartName string, chartVersion string, chartNamePath string, chartVersionPath string) error {
	//log.Printf("Updating chart version for '%s' to '%s'", chartName, chartVersion)
	namePathSegments := strings.Split(chartNamePath, ".")
	versionPathSegments := strings.Split(chartVersionPath, ".")

	// Find the map entry that has the target chartName
	found, err := j.findAndUpdate(j.Data, namePathSegments, versionPathSegments, chartName, chartVersion)
	if err != nil {
		log.Printf("Error during find and update: %v", err)
		return err
	}
	if !found {
		err := errors.New(fmt.Sprintf("chart name { %s } not found in the provided path { %v }", chartName, namePathSegments))
		log.Printf(err.Error())
		return err
	}

	//log.Printf("Successfully updated chart version for '%s'", chartName)
	return nil
}

// findAndUpdate recursively searches for the path and updates the version when the chart name is found
func (j *JSONDataWrapper) findAndUpdate(currentMap map[string]interface{}, namePathSegments, versionPathSegments []string, chartName, chartVersion string) (bool, error) {
	if len(namePathSegments) == 0 {
		err := errors.New("invalid chartName path")
		log.Printf("FAIL: %s", err)
		return false, err
	}

	currentPath := namePathSegments[0]
	if currentPath == "*" {
		for _, value := range currentMap {
			if subMap, ok := value.(map[string]interface{}); ok {
				// Check if the next segment after the wildcard matches the chart name
				if subMap[namePathSegments[1]] == chartName {
					//oldVersion, ok := j.findCurrentVersion(subMap, versionPathSegments)
					if !ok {
						log.Printf("FAIL: unable to find current version for '%s'", chartName)
						return false, errors.New("current version not found")
					}
					updated, err := j.updateVersion(subMap, versionPathSegments, chartName, chartVersion)
					if err != nil {
						//log.Printf("FAIL updating chart version for '%s' from '%s' to '%s': %v", chartName, oldVersion, chartVersion, err)
						return false, err
					}
					if updated {
						//log.Printf("DONE updating chart version for '%s' from '%s' to '%s'", chartName, oldVersion, chartVersion)
					}
					return updated, nil
				}
			}
		}
		return false, nil
	} else {
		if nextMap, ok := currentMap[currentPath].(map[string]interface{}); ok {
			return j.findAndUpdate(nextMap, namePathSegments[1:], versionPathSegments, chartName, chartVersion)
		}
	}

	log.Printf("FAIL: '%s' not found in the path", chartName)
	return false, nil
}

// findCurrentVersion finds the current version of the chart within the JSON map, based on the version path segments
func (j *JSONDataWrapper) findCurrentVersion(jsonData map[string]interface{}, pathSegments []string) (string, bool) {
	for i, segment := range pathSegments {
		if i == len(pathSegments)-1 {
			if version, exists := jsonData[segment]; exists {
				return version.(string), true
			}
			return "", false
		}

		if nextMap, ok := jsonData[segment].(map[string]interface{}); ok {
			jsonData = nextMap
		} else {
			return "", false
		}
	}
	return "", false
}

// updateVersion sets the chartVersion in the map based on the given path and logs the outcome
func (j *JSONDataWrapper) updateVersion(jsonData map[string]interface{}, pathSegments []string, chartName, newChartVersion string) (bool, error) {
	var oldVersion string
	updated := false
	// Traverse the JSON structure along the path segments to find where to update the version
	for i, segment := range pathSegments {
		if i == len(pathSegments)-1 {
			if currentVersion, exists := jsonData[segment]; exists {
				oldVersion = currentVersion.(string)
				jsonData[segment] = newChartVersion
				updated = true
				break
			} else {
				return false, errors.New("chartVersion path does not exist")
			}
		} else {
			if nextMap, ok := jsonData[segment].(map[string]interface{}); ok {
				jsonData = nextMap
			} else {
				return false, errors.New("invalid chartVersion path")
			}
		}
	}

	if updated {
		log.Printf("DONE updating chart version for '%s' from '%s' to '%s'", chartName, oldVersion, newChartVersion)
	} else {
		log.Printf("FAIL updating chart version for '%s' from '%s' to '%s'", chartName, oldVersion, newChartVersion)
	}
	return updated, nil
}

// WriteToFilesystem marshals the JSONDataWrapper's Data field to JSON and writes it to the specified file
func (j *JSONDataWrapper) WriteToFilesystem(filePath string) error {
	data, err := json.MarshalIndent(j.Data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
