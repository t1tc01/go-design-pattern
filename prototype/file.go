package prototype

import "fmt"

type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

// shallow copy, clone
func (f *File) Clone() Inode {
	return &File{Name: f.Name + "_clone"}
}
