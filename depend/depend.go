// Package depend provides an interface 
//
package depend

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var (
	dependmap = make(map[string]*DependNode)
	keyword   = []string{"all", "clean"}
)

// DependNode is a set of dependencies and associate recipes for a makefile
type DependNode struct {
	Name       string
	dependent  []string
	dependency []string
	recipe     []string
	DependList []*DependNode
  Run        chan bool
}

// ParseMake opens a makefile, and creates a list of the dependencies within the file
func ParseMake(filename string) map[string]*DependNode {
	// Open the makefile
	file, err := os.Open(filename)
	if err != nil {
	}
	defer file.Close()

	// build some regular expressions to identify lines in the file
	commentLine, err := regexp.Compile("#.*\n")
	dependencyLine, err := regexp.Compile("[A-Za-z]*[ ]*:.*\n")
	recipeLine, err := regexp.Compile("\t[^#].*\n")
	// TODO: Add in support for makefile variables

	// create an object to parse through the makefile
	scanner := bufio.NewReader(file)

	// begin reading lines from the makefile
	for line, err := scanner.ReadString('\n'); err != io.EOF; line, err = scanner.ReadString('\n') {
		switch {
		case commentLine.MatchString(line): // input is a comment so ignore
		case dependencyLine.MatchString(line):
			dependnode := new(DependNode)
      dependnode.Run = make(chan bool)

			//for a dependency the proceeding lines need to be checked for recipes
			dependencyArray := strings.Split(line, ":")

			// list over all of the dependent items and add to a node
			for _, d := range strings.Fields(dependencyArray[0]) {
				dependnode.dependent = append(dependnode.dependent, d)
			}

			// list over all of the dependencies and add to a node
			for _, d := range strings.Fields(dependencyArray[1]) {
				d = strings.TrimSpace(d)
				dependnode.dependency = append(dependnode.dependency, d)
			}

			//
			for line, err := scanner.ReadString('\n'); err != io.EOF && recipeLine.MatchString(line); line, err = scanner.ReadString('\n') {
				recipe := line
				dependnode.recipe = append(dependnode.recipe, recipe)
			}

			dependnode.Name = dependnode.dependent[0]
			dependmap[dependnode.Name] = dependnode

			// check if the dependency exists in the previous entries
Top:
			for _, d := range dependmap {
				for _, e := range d.dependency {
					for _, f := range dependnode.dependent {
						if e == f {
							d.DependList = append(d.DependList, dependnode)
							break Top
						}
					}
				}
			}
		}
	}
	return dependmap
}

// Make executes the recipes contained in the node after completing those below it and returns the number of times it executed
func (d *DependNode) Make(boolchan chan bool)  {

  forever := <-boolchan

  // execute the make
  doRecipe(*d)

  // inform the caller that you've finished
  boolchan <- forever

  // Wait to be called to execute again
  for {
    cond := <-d.Run
    if cond {
      doRecipe(*d)
    }
  }
}

// runs all the recipes if they need to be updated
func doRecipe(d DependNode) {

  // all the channels have to initialized separately
  arrayChan := make([]chan bool,len(d.DependList))
  for i := range arrayChan {
    arrayChan[i] = make(chan bool)
  }

  // if the node has other dependencies make those first
	for i,e := range d.DependList {
		go e.Make(arrayChan[i])
  }

  // send a signal to all of the dependencies
  for _,ac := range arrayChan {
    ac <- true
  }

  // wait for all of the dependencies to complete
  for _,ac := range arrayChan {
    <-ac
  }

  // get the most recently created dependent file
	minDependent := mostRecent(d.dependent)

	// get the most recently created dependency file
	minOtherTime := mostRecent(d.dependency)
	// once those have finished make this node
	if minDependent.Before(minOtherTime) || minDependent.IsZero() {
		for _, r := range d.recipe {
      cmd := exec.Command("sh","-c",r)
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(strings.TrimSpace(strings.Join(cmd.Args[2:]," ")))
		}
	}
}

// Returns the file created most recently
func mostRecent(files []string) time.Time {
	var mintime time.Time
	var temptime time.Time
	for _, e := range files {
		fi, err := os.Stat(e)
		if err != nil {
		} else {
			temptime = fi.ModTime()
		}
		if mintime.Before(temptime) {
			mintime = temptime
		}
	}
	return mintime
}

func contains(s string, strArray []string) bool {
	for _, word := range strArray {
		if word == s {
			return true
		}
	}
	return false
}
