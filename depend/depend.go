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
}

// Make executes the recipes contained in the node after completing those below it
func (d *DependNode) Make() {
}
