package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/leiwang008/utils"
)

var (
	help        bool
	source      string
	destination string
	verbose     bool
	checkdate   bool
	copy        bool
	overwrite   bool
)

func init() {
	flag.BoolVar(&help, "h", false, "get help messages")
	flag.BoolVar(&verbose, "v", false, "show verbose messages")
	flag.BoolVar(&checkdate, "ck", false, "check file's date and create subfolder according to it.")
	flag.BoolVar(&copy, "cp", false, "copy files if ture; move files if false")
	flag.BoolVar(&overwrite, "ov", false, "overwrite destination file if true; otherwise skip the file.")

	flag.StringVar(&source, "s", "", "the source folder")
	flag.StringVar(&destination, "d", "", "the destination folder")

}

func main() {
	debugmsg := utils.Debugmsg(true)

	fmt.Printf(debugmsg+"parameters %v\n", os.Args)

	flag.Parse()

	if help {
		usage()
		return
	}

	if utils.Empty(source) {
		log.Printf("The source cannot be empty!")
		usage()
		return
	}
	if utils.Empty(destination) {
		log.Printf("The destination cannot be empty!")
		usage()
		return
	}

	utils.Verbose = verbose
	err := utils.HandleFiles(source, destination, checkdate, copy, overwrite)

	if err != nil {
		log.Fatalf("Failed to handle file, due to %q\n", err)
	}

}

func usage() {
	// go run copy.go [-h] [-s source] [-d destination] [-v] [-ck] [-cp] [-ov]
	// --help, -h
	// --source, -s
	// --dest, -d
	// --verbose, -v
	// --checkdate, -ck
	// --copy, -cp
	// --overwrite, -ov
	fmt.Fprintf(os.Stdout, `copy version: copy/1.0.0
Usage: copy [-h] [-s source] [-d destination] [-v] [-ck] [-cp] [-ov]
Options:
`)
	flag.PrintDefaults()
}
