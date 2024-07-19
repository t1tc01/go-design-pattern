package prototype

/*
Shallow copy object
*/
type Inode interface {
	Print(string)
	Clone() Inode
}
