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

  /* Open the makefile */
  file,err := os.Open(*filename)
	if err != nil {
    return
  }
  defer file.Close()

  /* create an object to read through the makefile */
  scanner := bufio.NewReader(file)

	// build regular expressions to match the different cases that appear in the makefile
	comment,err := regexp.Compile("#.*\n")
	dependency,err:= regexp.Compile("[A-Za-z]*[ ]*:.*\n")
	recipe,err:= regexp.Compile("\t[^#].*\n")


  /* Begin reading in lines from the makefile */
  line,err := scanner.ReadString('\n')
  for err != io.EOF {
    /* compare to regular expressions */
    switch {
      case comment.MatchString(line):
      // comment so ignore
      case dependency.MatchString(line):
      // a dependency, check to see that it has a recipe on the next line
      // here we need to build a dependency object
        sourceArray := string.Split(line,":")
        string.Split(blank[0]," ")
        string.Split(blank[1]," ")
        line, err := scanner.ReadString('\n')
        for err != io.EOF || recipe.MatchString(line) {
          recipeArray := strings.SplitAfterN(line," ",2)
          fmt.Printf(recipeArray[0],recipeArray[1])
          line, err := scanner.ReadString('\n')
        }
    }
    line, err = scanner.ReadString('\n')
	}
}
