package info

import nomalfileinfo "github.com/yimkh/ostask/fileSystem/file/info"

//DirectoryNode is
type DirectoryNode struct {
	Tag  int
	Name string
	//0 nomal 1 directory
	Type  int
	Nfile *nomalfileinfo.NomalFile
	Dfile *DirectoryFile
}

//DirectoryFile is directory file
type DirectoryFile struct {
	DirectoryNodes []DirectoryNode
}
