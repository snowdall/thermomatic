// Package common implements utilities & functionality commonly consumed by the
// rest of the packages.
package common

import "errors"
import "os"
import "fmt"

// ErrNotImplemented is raised throughout the codebase of the challenge to
// denote implementations to be done by the candidate.
var ErrNotImplemented = errors.New("not implemented")

// This will output string to STDERR for reading on the prompt
func Err(input error) {
  os.Stdout.WriteString("ERROR: ")
  fmt.Fprintln(os.Stderr, input)
  os.Stdout.WriteString("/n")
}

// This will output string to STDOUT for reading on the prompt
func Out(input string) {
  os.Stdout.WriteString("OUTPUT: " + input + "\n")
}

