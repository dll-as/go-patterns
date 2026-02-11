package main

import "fmt"

// Component interface
type FileSystem interface {
	Print(prefix string)
}

// Leaf
type File struct {
	name string
}

func (f *File) Print(prefix string) {
	fmt.Println(prefix + f.name)
}

// Composite
type Folder struct {
	name     string
	children []FileSystem
}

func (f *Folder) Add(child FileSystem) {
	f.children = append(f.children, child)
}

func (f *Folder) Print(prefix string) {
	fmt.Println(prefix + f.name + "/")
	for _, child := range f.children {
		child.Print(prefix + "  ")
	}
}

func main() {
	// Create files and folders
	file1 := &File{name: "file1.txt"}
	file2 := &File{name: "file2.txt"}

	folder1 := &Folder{name: "folder1"}
	folder2 := &Folder{name: "folder2"}

	folder1.Add(file1)
	folder1.Add(file2)
	folder2.Add(folder1)
	folder2.Add(&File{name: "file3.txt"})

	// Print full structure
	folder2.Print("")
}
