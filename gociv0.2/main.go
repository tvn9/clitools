package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// declare command flags
	proj := flag.String("p", "", "Project directory")
	flag.Parse()

	// call the run function
	if err := run(*proj, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run
func run(proj string, out io.Writer) error {
	if proj == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)
	}

	pipeline := make([]step, 2)

	pipeline[0] = newStep(
		"go build",
		"go",
		"Go Build: SUCCESS",
		proj,
		[]string{
			"build", ".", "errors",
		},
	)
	for _, s := range pipeline {
		msg, err := s.execute()
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(out, msg)
		if err != nil {
			return err
		}
	}
	return nil
}
