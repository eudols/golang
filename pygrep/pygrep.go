package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func fileSearch(fullpath string, pattern string) {
	lineNumber := 0
	// fmt.Printf("I am in fileSearch for %s , and file pattern %s\n", fullpath, pattern)
	var regExp = regexp.MustCompile(pattern)
	f, err := os.Open(fullpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "*** Could not open %s\n", fullpath)
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		lineNumber++
		if regExp.MatchString(input.Text()) {
			fmt.Printf("%s:%d %s\n", fullpath, lineNumber, input.Text())
		}
	}
	f.Close()
}

func printFile(filepattern string, pattern string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		// fmt.Println("printFile filepattern ", filepattern)
		// fmt.Println("printFile pattern ", pattern)
		// fmt.Println("printFile path ", path)
		var regExp = regexp.MustCompile(filepattern)

		if regExp.MatchString(info.Name()) {
			// fmt.Printf("File %s matches\n", info.Name())
			fileSearch(path, pattern)
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
		fmt.Println("  (see \"https://golang.org/pkg/regexp/\" for details on regular expressions)")
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
	// fmt.Println("filepattern ", filepattern)
	// fmt.Println("pattern ", pattern)

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
