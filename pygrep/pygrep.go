package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func printFile(filepattern string, pattern string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		fmt.Println(filepattern)
		fmt.Println(pattern)
		fmt.Println(path)
		var regExp = regexp.MustCompile(filepattern)

		if regExp.MatchString(info.Name()) {
			fmt.Printf("File %s matches\n", info.Name())
		}
		return nil
	}
}

//
// func printFile(path string, info os.FileInfo, err error) error {
// 	if err != nil {
// 		log.Print(err)
// 		return nil
// 	}
// 	if info.IsDir() {
// 		fmt.Println("I am a directory, skip me")
// 		return nil
// 	}
// 	fmt.Print(info.Name())
// 	fmt.Println(path)
// 	return nil
// }

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

	log.SetFlags(log.Lshortfile)
	dir := "."
	err := filepath.Walk(dir, printFile(filepattern, pattern))
	if err != nil {
		log.Fatal(err)
	}
	// argsWithProg := os.Args
	// argsWithoutProg := os.Args[1:]
	//
	// fmt.Println(argsWithProg)
	// fmt.Println(argsWithoutProg)
}
