package main

import (
  "flag"
  "log"
  "os"
  "path/filepath"
  "regexp"
  "semver/semver"
)

func main() {
  var increaseBy, prefix, filenamePattern, fsPath string
  flag.StringVar(&increaseBy, "increase-semver-by-step", "", "Version to increase by")
  flag.StringVar(&prefix, "semver-prefix-string-regexp", "", "Regexp prefix for version string")
  flag.StringVar(&filenamePattern, "filename-pattern-regexp", "", "Filename regexp pattern")
  flag.StringVar(&fsPath, "filesystem-recursive-path", "./", "Filesystem path to start the search")
  flag.Parse()

  log.Printf("")
  log.Printf(`Regexp pattern for semver search in files: "%s"`, semver.CompileSemverIdentificationRegexpPattern(prefix))

  err := filepath.Walk(fsPath, func(path string, info os.FileInfo, err error) error {
    if err != nil {
      log.Printf("Error encountered while walking the filesystem: %v", err)
      return err
    }
    return processPath(path, filenamePattern, increaseBy, prefix)
  })

  if err != nil {
    log.Fatalf("Error processing filesystem: %v", err)
  }
}

func processPath(path string, filenamePattern, increaseBy, prefix string) error {
  info, err := os.Stat(path)
  if err != nil {
    log.Printf("Error getting file info for { %s }: %v", path, err)
    return err
  }

  matched, _ := regexp.MatchString(filenamePattern, info.Name())
  if !info.IsDir() && matched {
    log.Printf("Processing file: %s", path)

    content, err := os.ReadFile(path)
    if err != nil {
      log.Printf("Error reading file { %s }: %v", path, err)
      return err
    }

    newContent, err := semver.ProcessFileContent(string(content), increaseBy, prefix)
    if err != nil {
      log.Printf("Error processing file content for { %s }: %v", path, err)
      return err
    }

    if string(content) != newContent {
      err := os.WriteFile(path, []byte(newContent), 0644)
      if err != nil {
        log.Printf("Error writing to file { %s }: %v", path, err)
        return err
      }
      //log.Printf("Updated file: %s", path)
    } else {
      log.Printf("No changes for file: %s", path)
    }
  }
  return nil
}
