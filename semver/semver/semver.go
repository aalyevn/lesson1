package semver

import (
  "fmt"
  "log"
  "regexp"
  "strconv"
  "strings"
)

type SemVer struct {
  Major, Minor, Patch int
}

func ParseSemVer(version string) (SemVer, error) {
  parts := strings.Split(version, ".")
  if len(parts) != 3 {
    log.Printf("Error: Invalid semver format for input %s", version)
    return SemVer{}, fmt.Errorf("invalid semver format")
  }

  major, err := strconv.Atoi(parts[0])
  if err != nil {
    log.Printf("Error converting major version { %s } to integer: %v", parts[0], err)
    return SemVer{}, err
  }
  minor, err := strconv.Atoi(parts[1])
  if err != nil {
    log.Printf("Error converting minor version { %s } to integer: %v", parts[1], err)
    return SemVer{}, err
  }
  patch, err := strconv.Atoi(parts[2])
  if err != nil {
    log.Printf("Error converting patch version { %s } to integer: %v", parts[2], err)
    return SemVer{}, err
  }

  return SemVer{Major: major, Minor: minor, Patch: patch}, nil
}

func (s *SemVer) IncrementBy(step SemVer) {
  s.Major += step.Major
  s.Minor += step.Minor
  s.Patch += step.Patch
}

func (s SemVer) String() string {
  return fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
}

func CompileSemverIdentificationRegexpPattern(prefix string) string {
  return prefix + `.*` + `(\d+\.\d+\.\d+)`
}

func ProcessFileContent(content, increaseBy, prefix string) (string, error) {
  step, err := ParseSemVer(increaseBy)
  if err != nil {
    log.Printf("Error parsing increaseBy version { %s }: %v", increaseBy, err)
    return content, err
  }

  regex := regexp.MustCompile(CompileSemverIdentificationRegexpPattern(prefix))

  newContent := regex.ReplaceAllStringFunc(content, func(s string) string {
    match := regex.FindStringSubmatch(s)
    if len(match) < 2 {
      log.Printf("No semver match found for string { %s }", s)
      return s
    }

    oldVersion, err := ParseSemVer(match[1])
    if err != nil {
      log.Printf("Error parsing version from match { %s }: %v", match[1], err)
      return s
    }

    oldVersionString := oldVersion.String() // store the old version string for logging
    oldVersion.IncrementBy(step)
    newVersionString := oldVersion.String() // get the new version string for logging

    log.Printf("Adjusting version from { %s } to { %s } by step { %s }", oldVersionString, newVersionString, increaseBy)

    return strings.Replace(s, match[1], newVersionString, 1)
  })

  return newContent, nil
}
