# Release Harvester

Release Harvester is a utility for updating chart versions in JSON files based on the contents found within Chart.yaml files throughout a given directory structure.

## Features

- Search through directories for Chart.yaml files.
- Extract chart version information and maintain a dictionary of the latest versions.
- Update chart versions in a specified JSON file based on a provided JSON path.
- A flexible CLI tool suitable for automating in continuous integration workflows.

## Prerequisites

- Golang (ensure you have a compatible version with the project, preferably the latest stable release)

## Installation

Clone the repository and navigate to the project directory:

```shell
git clone [REPOSITORY_URL]
cd [LOCAL_DIRECTORY]
```

## Build
To build the application, run the following command in the project directory:
```shell
go build -o release_harvester
```


## Usage
To use the tool, execute the binary with the required flags:
```shell
./release_harvester --destination-json [PATH_TO_JSON] --chart-name-json-path [JSON_PATH] --chart-version-subpath [VERSION_SUBPATH] --charts-recursive-location [CHARTS_DIRECTORY]
```

### Flags:
* `--destination-json`: Path to the destination JSON file where the chart versions will be updated.
* `--chart-name-json-path`: The JSON path expression to locate the chart name within the destination JSON file.
* `--chart-version-subpath`: The subpath to locate and update the chart version in the destination JSON file.
* `--charts-recursive-location`: The filesystem path to recursively search for Chart.yaml files.

## Examples
To update chart versions in test.json by searching through all directories starting with the current one:
```shell
./release_harvester --destination-json "test.json" --chart-name-json-path "mocks.*.chartName" --chart-version-subpath "chartVer" --charts-recursive-location ./
```

## Contributing
Contributions are welcome. If you have suggestions for improvement or have uncovered an issue, please open an issue first to discuss what you would like to change or report the problem.

## License
This project is open-sourced under the MIT License. See the LICENSE file for more details.



