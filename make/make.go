package main

import (
  //"fmt"
  //"os"
  "flag"
  //"bufio"
  //"io"
  //"strings"
  //"regexp"
  "log"
  "gomake/depend"
  "github.com/howeyc/fsnotify"
)

// The command line argument flags for the gomake program
var usage = flag.Bool("help",false,"Prints this help message.")
var filename = flag.String("f","Makefile","Read FILE as a makefile.")
var version = flag.Float64("v",1.0,"Print the version number of gomake and exit")
var debug = flag.Bool("d",false,"Print lots of debugging information")
var forever = flag.Bool("forever",false,"Run the program continously")

func main() {
  name := "all"
  boolchan := make(chan bool)
	flag.Parse()
	if flag.NArg() > 0 {
    name = flag.Arg(0)
  }
  dependmap := depend.ParseMake(*filename)
  dmap := dependmap[name]
  go dmap.Make(boolchan)
  boolchan <- *forever
  *forever = <-boolchan

  // if the forever option has been chosen, loop forever
  if *forever {
    // initiate a watcher
    watcher,err := fsnotify.NewWatcher()
    if err != nil {
      log.Fatal(err)
    }

    go func() {
      for {
        select {
          case ev := <-watcher.Event:
            for _,dmap := range dependmap {
              if ev.Name == "./" + dmap.Name {
                dmap.Run <- true
              }
            }
          case err := <-watcher.Error:
            log.Fatal(err)
        }
      }
    }()

    // Set the path to watch
    err = watcher.Watch(".")
    if err != nil {
      log.Fatal(err)
    }

    for {
      <-boolchan
    }
  }
}
