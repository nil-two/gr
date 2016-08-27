package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ogier/pflag"
)

var (
	cmdName    = "gr"
	cmdVersion = "0.0.0"

	flagset      = pflag.NewFlagSet(cmdName, pflag.ContinueOnError)
	findNearArg  = flagset.StringP("near", "n", "", "")
	findOfArg    = flagset.StringP("of", "o", "", "")
	findRangeArg = flagset.StringP("range", "r", "", "")
	firstNumber  = flagset.Float64P("first", "f", 3.0, "")
	showDecimal  = flagset.BoolP("decimal", "d", false, "")
	isHelp       = flagset.BoolP("help", "", false, "")
	isVersion    = flagset.BoolP("version", "", false, "")
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s OPTION...
Find the golden number you want.

Predicates:
  -n, --near=N      find numbers near N
  -o, --of=N        find numbers of N
  -r, --range=N-M   find numbers between N and M

Miscellaneous:
  -f, --first=N     set first golden number for N (default: 3.0)
  -d, --decimal     show number as a decimal
      --help        display this help and exit
      --version     display version information and exit
`[1:], cmdName)
}

func printVersion() {
	fmt.Fprintf(os.Stderr, "%s\n", cmdVersion)
}

func printErr(err interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", cmdName, err)
}

func guideToHelp() {
	fmt.Fprintf(os.Stderr, "Try '%s --help' for more information.\n", cmdName)
}

func parsePredicate() (p Predicate, err error) {
	nPredicates := 0
	flagset.Visit(func(f *pflag.Flag) {
		switch f.Name {
		case "near", "of", "range":
			nPredicates++
		}
	})
	if nPredicates < 1 {
		return nil, fmt.Errorf("no spcify predicate")
	}
	if nPredicates > 1 {
		return nil, fmt.Errorf("only one type of predicate may be specified")
	}

	switch {
	case *findNearArg != "":
		return NewFindNear(*findNearArg)
	case *findOfArg != "":
		return NewFindOf(*findOfArg)
	case *findRangeArg != "":
		return NewFindRange(*findRangeArg)
	default:
		return nil, fmt.Errorf("reached unreachable case")
	}
}

func _main() int {
	flagset.SetOutput(ioutil.Discard)
	if err := flagset.Parse(os.Args[1:]); err != nil {
		printErr(err)
		guideToHelp()
		return 2
	}
	if *isHelp {
		printUsage()
		return 0
	}
	if *isVersion {
		printVersion()
		return 0
	}

	p, err := parsePredicate()
	if err != nil {
		printErr(err)
		guideToHelp()
		return 2
	}

	p.Init(*firstNumber, *showDecimal)
	if err = p.Do(os.Stdout); err != nil {
		printErr(err)
		return 1
	}
	return 0
}

func main() {
	e := _main()
	os.Exit(e)
}
