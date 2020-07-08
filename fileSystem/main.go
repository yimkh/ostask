package main

import (
	"fmt"

	"github.com/yimkh/ostask/fileSystem/command"
	dinfo "github.com/yimkh/ostask/fileSystem/dirfile/info"
	ninfo "github.com/yimkh/ostask/fileSystem/file/info"
)

func main() {
	user := dinfo.DirectoryNode{
		Tag:   1,
		Name:  "user",
		Type:  1,
		Dfile: &dinfo.DirectoryFile{},
		Nfile: &ninfo.NomalFile{},
	}

	app := dinfo.DirectoryNode{
		Tag:   1,
		Name:  "app",
		Type:  1,
		Dfile: &dinfo.DirectoryFile{},
		Nfile: &ninfo.NomalFile{},
	}
	download := dinfo.DirectoryNode{
		Tag:   1,
		Name:  "download",
		Type:  1,
		Dfile: &dinfo.DirectoryFile{},
		Nfile: &ninfo.NomalFile{},
	}

	user.Dfile.DirectoryNodes = append(user.Dfile.DirectoryNodes, app)
	user.Dfile.DirectoryNodes = append(user.Dfile.DirectoryNodes, download)

	nowNode := dinfo.DirectoryNode{}
	nowNode = user

	var inputCommand string
	var parameter string
	var size int

label:
	for {
		fmt.Println("please input your command(EXIT CREwill exit):")
		fmt.Scanln(&inputCommand, &parameter)
		switch inputCommand {
		case "CD":
			nowNode = command.CD(nowNode, parameter)
			fmt.Printf("%s+~\n", nowNode.Name)
		case "CREATE":
			fmt.Println("please input the size of file")
			fmt.Scanln(&size)
			nowNode = command.CREATE(nowNode, parameter, 202001, size)
			printn(nowNode.Nfile.FileNodes)
		case "DEL":
			nowNode = command.DEL(nowNode, parameter)
			printn(nowNode.Nfile.FileNodes)
		case "LSALL":
			command.LSALL(nowNode)
		case "MD":
			nowNode = command.MD(nowNode, parameter)
			printd(nowNode.Dfile.DirectoryNodes)
		case "RD":
			nowNode = command.RD(nowNode, parameter)
			printd(nowNode.Dfile.DirectoryNodes)
		case "EXIT":
			break label
		default:
			fmt.Println("keyword is error, please input again")
		}
	}
}

func printn(FileNodes []ninfo.FileNode) {
	fmt.Print("NOMAL File:  ")
	for _, v := range FileNodes {
		fmt.Print(v.Name + "  ")
	}
	fmt.Println()
}

func printd(DirectoryNodes []dinfo.DirectoryNode) {
	fmt.Print("Directory File:  ")
	for _, v := range DirectoryNodes {
		fmt.Print(v.Name + "  ")
	}
	fmt.Println()
}
