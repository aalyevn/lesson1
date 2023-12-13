package main

import (
	"flag"
	"fmt"
	"json_editor/chartparser"
	"json_editor/jsonwrapper"
	"os"
)

func main() {
	// Parse command-line arguments
	destinationJSON, chartNameJSONPath, chartVersionSubpath, chartsRecursiveLocation := parseFlags()

	// Setup JSON editor with provided JSON file
	jsonDataWrapper := setupJSONEditor(destinationJSON)

	// Create and populate a dictionary of chart versions from specified location
	dictionary := createAndPopulateDictionary(chartsRecursiveLocation)

	// Update chart versions in JSON data using the dictionary
	updateChartVersions(jsonDataWrapper, dictionary, chartNameJSONPath, chartVersionSubpath)

	// Save changes to the JSON file
	saveJSONEdits(jsonDataWrapper, destinationJSON)

	fmt.Println("Chart versions updated successfully from the dictionary.")
}

func parseFlags() (destinationJSON, chartNameJSONPath, chartVersionSubpath, chartsRecursiveLocation string) {
	flag.StringVar(&destinationJSON, "destination-json", "test.json", "The destination JSON file path")
	flag.StringVar(&chartNameJSONPath, "chart-name-json-path", "mocks.*.chartName", "The JSON path to locate the chart name in destination JSON file")
	flag.StringVar(&chartVersionSubpath, "chart-version-subpath", "chartVer", "The subpath to locate and update the chart version in destination JSON file (inside the entry, which corresponds to the identified chart-name)")
	flag.StringVar(&chartsRecursiveLocation, "charts-recursive-location", ".", "The location to recursively search for charts (searching for Chart.yaml)")
	flag.Parse()
	return
}

func setupJSONEditor(filePath string) *jsonwrapper.JSONDataWrapper {
	jsonDataWrapper, err := jsonwrapper.New(filePath)
	if err != nil {
		exitWithError(fmt.Sprintf("Initializing JSON data wrapper failed: %v", err))
	}
	return jsonDataWrapper
}

func createAndPopulateDictionary(recursiveLocation string) *chartparser.Dictionary {
	dictionary := chartparser.NewDictionary()
	if err := dictionary.PopulateFromWalk(recursiveLocation); err != nil {
		exitWithError(fmt.Sprintf("Populating chart dictionary failed: %v", err))
	}
	return dictionary
}

func updateChartVersions(jsonDataWrapper *jsonwrapper.JSONDataWrapper, dictionary *chartparser.Dictionary, chartNamePath, chartVersionPath string) {
	if err := jsonDataWrapper.UpdateChartVersionFromDictionary(dictionary.Charts, chartNamePath, chartVersionPath); err != nil {
		exitWithError(fmt.Sprintf("Updating chart versions failed: %v", err))
	}
}

func saveJSONEdits(jsonDataWrapper *jsonwrapper.JSONDataWrapper, destinationJSON string) {
	if err := jsonDataWrapper.WriteToFilesystem(destinationJSON); err != nil {
		exitWithError(fmt.Sprintf("Saving JSON edits failed: %v", err))
	}
}

func exitWithError(message string) {
	fmt.Println(message)
	os.Exit(1)
}
