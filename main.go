package main

import (
	factorymethod "example/go-design-pattern/factory_method"
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

func FactoryMethodExample() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factorymethod.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}

func getGun(gunType string) (factorymethod.IGun, error) {
	if gunType == "ak47" {
		return factorymethod.NewAk47(), nil
	}
	if gunType == "musket" {
		return factorymethod.NewMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}

// main
func main() {
	fmt.Println("Run....")
}
