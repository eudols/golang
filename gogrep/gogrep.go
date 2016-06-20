package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	pattern := ""
	filepattern := `\.(h|c|cc|cpp)$` // default to source files
	var argM, argC, argP bool

	flag.Usage = func() {
		fmt.Println("Usage: cgrep [-mp] [-f <file regexp>] <regexp>")
		fmt.Println("  -h: this help")
		fmt.Println("  -m: search in make files")
		fmt.Println("  -p: search in python files")
		fmt.Println("  -f: use an alternate file regular expression")
		fmt.Println("  (see \"man perlre\" for perl regular expressions)")
	}

	flag.BoolVar(&argC, "c", false, "")
	flag.BoolVar(&argM, "m", false, "")
	flag.BoolVar(&argP, "p", false, "")
	flag.StringVar(&filepattern, "f", `\.(h|c|cc|cpp)$`, "")
	flag.Parse()

	if argC {
		filepattern = `^Cons(cript|struct)$`
	} else if argM {
		filepattern = `([Mm]ake[^~]*|\.mk)$`
	} else if argP {
		filepattern = `\.py$`
	}
	// if no arguments, print usage and exit
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	pattern = flag.Arg(0)
	fmt.Println(filepattern)
	fmt.Println(pattern)

	// argsWithProg := os.Args
	// argsWithoutProg := os.Args[1:]
	//
	// fmt.Println(argsWithProg)
	// fmt.Println(argsWithoutProg)
}
