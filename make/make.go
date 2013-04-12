package main

import (
  "fmt"
  "os"
  "flag"
  "bufio"
  "io"
  "strings"
  "regexp"
  "gomake/depend"
)

// The command line argument flags for the gomake program
var usage = flag.Bool("help",false,"Prints this help message.")
var filename = flag.String("f","Makefile","Read FILE as a makefile.")
var version = flag.Float64("v",1.0,"Print the version number of gomake and exit")
var debug = flag.Bool("d",false,"Print lots of debugging information")

func main() {
	flag.Parse()
	if flag.NArg() > 0 {
  }

  listofdependencies := Depend.ParseMake(filename)

	}
}
