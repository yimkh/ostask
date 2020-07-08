package command

import (
	"fmt"

	dinfo "github.com/yimkh/ostask/fileSystem/dirfile/info"
	ninfo "github.com/yimkh/ostask/fileSystem/file/info"
)

//CD is ender a DirectoryNode
func CD(nowNode dinfo.DirectoryNode, name string) dinfo.DirectoryNode {
	flag := len(nowNode.Dfile.DirectoryNodes)

	for i := 0; i < len(nowNode.Dfile.DirectoryNodes); i++ {
		if nowNode.Dfile.DirectoryNodes[i].Name == name {
			nowNode = nowNode.Dfile.DirectoryNodes[i]
			return nowNode
		}

		flag--
	}

	if flag == 0 {
		fmt.Printf("%s is not exist", name)
	}

	return nowNode
}

//CREATE is to create
func CREATE(nowNode dinfo.DirectoryNode, name string, thisType int, size int) dinfo.DirectoryNode {
	nowNode.Nfile.FileNodes = append(nowNode.Nfile.FileNodes, ninfo.FileNode{
		Name: name,
		Type: thisType,
		Size: size,
	})

	return nowNode
}

//DEL is to delete nomal file
func DEL(nowNode dinfo.DirectoryNode, name string) dinfo.DirectoryNode {
	for i := 0; i < len(nowNode.Nfile.FileNodes); i++ {
		if nowNode.Nfile.FileNodes[i].Name == name {
			nowNode.Nfile.FileNodes = append(nowNode.Nfile.FileNodes[:i], nowNode.Nfile.FileNodes[i+1:]...)

			return nowNode
		}
	}

	fmt.Printf("this %s is not exist\n", name)

	return nowNode
}

//LSALL is to show all diretory
func LSALL(nowNode dinfo.DirectoryNode) {
	fmt.Print("DirectoryFile:------")
	for _, v := range nowNode.Dfile.DirectoryNodes {
		fmt.Print(v.Name + "  ")
	}
	fmt.Println()
	fmt.Print("NomalFile:------")
	for _, v := range nowNode.Nfile.FileNodes {
		fmt.Print(v.Name + "  ")
	}
	fmt.Println()
}

//MD is to make a directory
func MD(nowNode dinfo.DirectoryNode, name string) dinfo.DirectoryNode {
	for i := 0; i < len(nowNode.Dfile.DirectoryNodes); i++ {
		if nowNode.Dfile.DirectoryNodes[i].Name == name {
			fmt.Printf("make directory error, already have %s\n", name)
			return nowNode
		}
	}

	nowNode.Dfile.DirectoryNodes = append(nowNode.Dfile.DirectoryNodes, dinfo.DirectoryNode{
		Tag:   1,
		Name:  name,
		Type:  1,
		Dfile: &dinfo.DirectoryFile{},
		Nfile: &ninfo.NomalFile{},
	})

	return nowNode
}

//RD is to remove directory
func RD(nowNode dinfo.DirectoryNode, name string) dinfo.DirectoryNode {
	for i := 0; i < len(nowNode.Dfile.DirectoryNodes); i++ {
		if nowNode.Dfile.DirectoryNodes[i].Name == name {
			if nowNode.Dfile.DirectoryNodes[i].Dfile == nil {
				//remove ........
				nowNode.Dfile.DirectoryNodes = append(nowNode.Dfile.DirectoryNodes[:i], nowNode.Dfile.DirectoryNodes[i+1:]...)

				return nowNode
			}

			fmt.Println("this directiry is not empty, input * to remove all")
			nowNode.Dfile.DirectoryNodes = append(nowNode.Dfile.DirectoryNodes[:i], nowNode.Dfile.DirectoryNodes[i+1:]...)

			return nowNode
		}
	}

	fmt.Printf("this %s is not exist\n", name)

	return nowNode
}
