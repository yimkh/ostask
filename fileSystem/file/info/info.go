package info

//FileNode is File Node
type FileNode struct {
	Name string
	//202001, 202002...
	Type int
	Size int
}

//NomalFile is Nomal File
type NomalFile struct {
	FileNodes []FileNode
}
