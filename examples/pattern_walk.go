// examples/pattern_walk.go

package main

import (
	clasp "github.com/synesissoftware/CLASP.Go"
	shwild "github.com/synesissoftware/shwild.Go"

	"fmt"
	"os"
	"path/filepath"
)

const (
	ProgramVersion = "0.0.1"
)

type ProcessFlag int64

const (
	ProcessFlag_ShowHidden = 1 << iota
)

func main() {

	// Specify aliases, parse, and checking standard flags

	flag_ShowHidden := clasp.Flag("--show-hidden").SetAlias("-h").SetHelp("includes hidden files in the search results")

	specifications := []clasp.Specification{

		flag_ShowHidden,

		clasp.HelpFlag(),
		clasp.VersionFlag(),
	}

	args := clasp.Parse(os.Args, clasp.ParseParams{Specifications: specifications})

	if args.FlagIsSpecified(clasp.HelpFlag()) {

		clasp.ShowUsage(specifications, clasp.UsageParams{

			Version:      ProgramVersion,
			InfoLines:    []string{"shwild.Go Examples", ":version:", "Walks the given root and all its subdirectories and reports all the files found that match the specified pattern(s)", ""},
			ValuesString: "{ <root> | . } <pattern-1> [... <pattern-N>]",
		})
	}

	if args.FlagIsSpecified(clasp.VersionFlag()) {

		clasp.ShowVersion(specifications, clasp.UsageParams{Version: ProgramVersion})
	}

	// Program-specific processing of flags/options

	var flags ProcessFlag

	if args.FlagIsSpecified(flag_ShowHidden) {

		flags |= ProcessFlag_ShowHidden
	}

	// Processing values

	if len(args.Values) < 2 {

		fmt.Fprintf(os.Stderr, "%s: Must specify directory and one or more pattern; use --help for usage\n", args.ProgramName)

		os.Exit(1)
	} else {

		// gather the remaining values and convert them into
		// compiled-patterns

		var directory string
		var patterns []shwild.CompiledPattern

		directory = args.Values[0].Value

		for _, value := range args.Values[1:] {

			pattern := value.Value
			cp, err := shwild.Compile(pattern)
			if err != nil {

				fmt.Fprintf(os.Stderr, "%s: invalid pattern '%s': %v\n", args.ProgramName, pattern, err)

				os.Exit(1)
			} else {

				patterns = append(patterns, cp)
			}
		}

		process(directory, patterns, flags, args.ProgramName)
	}
}

func process(directory string, patterns []shwild.CompiledPattern, flags ProcessFlag, program_name string) {

	err := filepath.Walk(directory, func(path string, fi os.FileInfo, err error) error {

		if fi.IsDir() {

			return nil
		}

		if 0 == (ProcessFlag_ShowHidden&flags) && is_hidden(path, fi) {

			return nil
		}

		// see if any of the patterns are matched

		match_any := false

		for _, p := range patterns {

			matched, err := p.Match(fi.Name())

			if matched && err == nil {

				match_any = true

				break
			}
		}

		if !match_any {

			return nil
		}

		fmt.Fprintf(os.Stdout, "found '%s' %d bytes\n", path, fi.Size())

		return nil
	})

	if err != nil {

		fmt.Fprintf(os.Stderr, "%s: search of '%s' failed: %v\n", program_name, directory, err)
	}
}

func is_hidden(path string, fi os.FileInfo) bool {

	if '.' == path[0] {

		return true
	}

	return '.' == fi.Name()[0]
}
