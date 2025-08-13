# shwild.Go Example - **pattern_walk**

## Summary

An example using **shwild.Go**'s compiled-patterns and the standard library's ```filepath.Walk()``` to do a recursive filtered search of a directory. The directory and patterns are specified as values on the command-line; matching behaviour is moderated by flags/options.

## Source

```Go
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

    ProgramVersion  =   "0.0.1"
)

type ProcessFlag int64

const (

    ProcessFlag_ShowHidden          =   1 << iota
)

func main() {

    // Specify aliases, parse, and checking standard flags

    flag_ShowHidden := clasp.Alias{ clasp.Flag, "--show-hidden", []string{ "-h" }, "includes hidden files in the search results", nil, 0 }

    aliases := []clasp.Alias {

        flag_ShowHidden,

        clasp.HelpFlag(),
        clasp.VersionFlag(),
    }

    args := clasp.Parse(os.Args, clasp.ParseParams{ Aliases: aliases })

    if args.FlagIsSpecified(clasp.HelpFlag()) {

        clasp.ShowUsage(aliases, clasp.UsageParams{

            Version: ProgramVersion,
            InfoLines: []string { "shwild.Go Examples", ":version:", "Walks the given root and all its subdirectories and reports all the files found that match the specified pattern(s)", "" },
            ValuesString: "{ <root> | . } <pattern-1> [... <pattern-N>]",
        })
    }

    if args.FlagIsSpecified(clasp.VersionFlag()) {

        clasp.ShowVersion(aliases, clasp.UsageParams{ Version: ProgramVersion })
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

        var directory   string
        var patterns    []shwild.CompiledPattern

        directory   =   args.Values[0].Value

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

    err := filepath.Walk(directory, func (path string, fi os.FileInfo, err error) error {

        if fi.IsDir() {

            return nil
        }

        if 0 == (ProcessFlag_ShowHidden & flags) && is_hidden(path, fi) {

            return nil
        }

        // see if any of the patterns are matched

        match_any   :=  false

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
```

## Usage

### No arguments

If executed with no arguments

```
$ go run examples/pattern_walk.go
```

it gives the output:

```
pattern_walk: Must specify directory and one or more pattern; use --help for usage
```

with an exit code of 1

### Show usage

If executed with the arguments

```
$ go run examples/pattern_walk.go --help
```

it gives the output:

```
shwild.Go Examples
pattern_walk 0.0.1
Walks the given root and all its subdirectories and reports all the files found that match the specified pattern(s)

USAGE: pattern_walk [ ... flags and options ... ] { <root> | . } <pattern-1> [... <pattern-N>]

flags/options:

	-h
	--show-hidden
		includes hidden files in the search results

	--help
		Shows this helps and exits

	--version
		Shows version information and exits
```

with an exit code of 1

### Searching in the project root

If executed with the arguments

```
$ go run examples/pattern_walk.go -h . '.*' '[ER]*' '*m*.go'
```

it gives the output:

```
found '.gitignore' 283 bytes
found 'EXAMPLES.md' 332 bytes
found 'README.md' 1033 bytes
found 'compiled_pattern_test.go' 13369 bytes
found 'examples/.pattern_walk.go.swp' 12288 bytes
found 'match_test.go' 7610 bytes
found 'matchers.go' 7082 bytes
```
