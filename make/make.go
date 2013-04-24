package main

import (
  "fmt"
  //"os"
  "flag"
  //"bufio"
  //"io"
  //"strings"
  //"regexp"
  "gomake/depend"
)

// The command line argument flags for the gomake program
var usage = flag.Bool("help",false,"Prints this help message.")
var filename = flag.String("f","Makefile","Read FILE as a makefile.")
var version = flag.Float64("v",1.0,"Print the version number of gomake and exit")
var debug = flag.Bool("d",false,"Print lots of debugging information")

func main() {
  name := "all"
	flag.Parse()
	if flag.NArg() > 0 {
    name = flag.Arg(0)
  }
  dependmap := depend.ParseMake(*filename)
  dmap := dependmap[name]
  x := dmap.Make(0)
  fmt.Println(dmap.Name)

  if x < 1 {
    fmt.Println("Nothing to make for",name)
  }
}
