# SemVer Increment CLI Tool
This CLI tool provides an easy way to recursively increment the semantic version (SemVer) in files matching a given filename pattern.

## Features
* Recursively search through directories for files matching a specified pattern.
* Increment the semantic version in files based on a provided step.
* Matches files by regex pattern and replaces version strings with the incremented value.
* 
## Prerequisites
* Golang (tested with version 1.17.5)

## Installation
Clone the repository and navigate to the project directory:
```shell
git clone git@github.com:DawnBreather/semver.git
cd ./semver
```

Run unit tests:
```shell
go test .
```


Build the application:
```shell
go build -o semver ./main/main.go
```

## Usage
```shell
./semver --increase-semver-by-step [INCREMENT] --semver-prefix-string-regexp [PREFIX] --filename-pattern-regexp [PATTERN] --filesystem-recursive-path [PATH]
```

## Parameters:
* `--increase-semver-by-step`: Define how much the version should be incremented. Example: "0.0.1" (increases patch version by 1), "1.0.1" (increases major and patch versions by 1), "0.0.4" (increases patch version by 4)
* `--semver-prefix-string-regexp`: Regex prefix to match the version string in the file. Example: "^version:"
* `--filename-pattern-regexp`: Regex pattern to match filenames. Example: "Chart.yaml"
* `--filesystem-recursive-path`: The filesystem path to start the search. Defaults to "./"

## Examples
To increment the patch version in all Chart.yaml files recursively:
```shell
./semver --increase-version-by-step 0.0.1 --string-regexp-prefix "version:" --filename-regexp-pattern "Chart.yaml" --recursive-filesystem-path ./
```

## Testing
Run the unit tests:
```shell
go test .
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)

