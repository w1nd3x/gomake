// Package depend provides an interface 
//
//
//
//
//
//
package depend

import (
  "os"
  "os/exec"
  "bufio"
  "io"
  "strings"
  "regexp"
)

// DependNode is a set of dependencies and associate recipes for a makefile
type DependNode struct {
  dependent FileInfo
  dependency FileInfo
  recipe []os.Cmd
  associatedDependencies []*DependNode
}

// ParseMake opens a makefile, and creates a list of the dependencies within the file
func ParseMake(filename string) []*DependNode {
  // Open the makefile
  file, err := os.Open(filename)
  if err != nil {

  }
  defer file.Close()

  // build some regular expressions to identify lines in the file
	commentLine,err := regexp.Compile("#.*\n")
	dependencyLine,err:= regexp.Compile("[A-Za-z]*[ ]*:.*\n")
	recipeLine,err:= regexp.Compile("\t[^#].*\n")
  // TODO: Add in support for makefile variables

  // the slice of nodes to return
  dependslice []*DependNode

  // create an object to parse through the makefile
  scanner := bufio.NewReader(file)

  // begin reading lines from the makefile
  for line, err := scanner.ReadString('\n'); err != io.EOF {
    switch {
      case commentLine.MatchString(line):  // input is a comment so ignore
      case dependencyLine.MatchString(line):
        dependnode := new DependNode
        // for a dependency the proceeding lines need to be checked for recipes
        dependencyArray := string.Split(line,":")
        for line, err := scanner.ReadString('\n'); err != io.EOF && recipeLine.MatchString(line) {
          recipeArray := strings.SplitAfterN(line," ",2)
          recipe := exec.Command(recipeArray[0], recipeArray[1])
          dependnode.recipe = append(dependnode.recipe, recipe)
        }

  }

}

// Make executes the recipes contained in the node after completing those below it
func (d *DependNode) Make() {
}
