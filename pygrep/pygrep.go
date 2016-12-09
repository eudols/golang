package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

// Takes a channel that receives string channels. Each matching file
// will have its own string channel. This function will iterate over
// each one in order
func collectLines(pipe chan chan string) {
	defer wg.Done()
	for a := range pipe {
		for str := range a {
			fmt.Println(str)
		}
	}
}

func fileSearch(fullpath string, pattern string, lines chan string) {
	defer wg.Done()
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
			//fmt.Printf("%s:%d %s\n", fullpath, lineNumber, input.Text())
			lines <- (fullpath + ":" + strconv.Itoa(lineNumber) + " " + input.Text())
		}
	}
	close(lines)
	f.Close()
}

func printFile(filepattern string, pattern string, a chan chan string) filepath.WalkFunc {
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
			wg.Add(1)
			lines := make(chan string, 100)
			a <- lines
			go fileSearch(path, pattern, lines)
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

	wg.Add(1)
	a := make(chan chan string)
	go collectLines(a)
	log.SetFlags(log.Lshortfile)
	dir := "."
	//dir := "/neteng/edols/supportassist/AR-9289-ETL-replacement/ngos/serv-supportassist/lib"
	err := filepath.Walk(dir, printFile(filepattern, pattern, a))
	if err != nil {
		log.Fatal(err)
	}
	wg.Done()
	wg.Wait()
	wg.Add(1)
	close(a)
	wg.Wait()
	// argsWithProg := os.Args
	// argsWithoutProg := os.Args[1:]
	//
	// fmt.Println(argsWithProg)
	// fmt.Println(argsWithoutProg)
}
