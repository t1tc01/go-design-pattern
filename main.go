package main

import (
	"example/go-design-pattern/prototype"
	"example/go-design-pattern/singleton"
	"fmt"
)

// Singleton
func SingletonExample() {
	for i := 0; i < 30; i++ {
		go singleton.GetInstancev2()
	}
	fmt.Scanln()

}

// Prototype
func PrototypeExample() {

	file1 := &prototype.File{Name: "File1"}
	file2 := &prototype.File{Name: "File2"}
	file3 := &prototype.File{Name: "File3"}

	folder1 := &prototype.Folder{
		Children: []prototype.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &prototype.Folder{
		Children: []prototype.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}

// main
func main() {
	fmt.Println("Run....")
}
