package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	pattern := ""
	filepattern := `\.py$` // default to Python source files
	var argC, argM, argJ bool

	flag.Usage = func() {
		fmt.Println("Usage: pygrep [-mcj] [-f <file regexp>] <regexp>")
		fmt.Println("  -h: this help")
		fmt.Println("  -c: search in C files")
		fmt.Println("  -m: search in make files")
		fmt.Println("  -j: search in javascript files")
		fmt.Println("  -f: use an alternate file regular expression")
		fmt.Println("  (see \"man perlre\" for perl regular expressions)")
	}

	flag.BoolVar(&argC, "c", false, "")
	flag.BoolVar(&argM, "m", false, "")
	flag.BoolVar(&argJ, "j", false, "")
	flag.StringVar(&filepattern, "f", `\.py$`, "")
	flag.Parse()

	if argC {
		filepattern = `\.(h|c|cc|cpp)$`
	} else if argM {
		filepattern = `([Mm]ake[^~]*|\.mk)$`
	} else if argJ {
		filepattern = `\.js$`
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
