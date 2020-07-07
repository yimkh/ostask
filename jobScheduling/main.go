package main

import (
	"fmt"

	hrrn "github.com/yimkh/ostask/jobScheduling/arithmatic"
	file "github.com/yimkh/ostask/jobScheduling/fileInfo"
)

func main() {
	var (
		Name    string
		GetTime int
		RunTime int
		Size    int
		Tape    int
		Printer int

		flag int
	)

	files := []file.File{}

	fmt.Println("input 1 to begin inputing tasks")
	fmt.Scan(&flag)

	for {
		if flag == 1 {
			fmt.Println("Please input task name: ")
			fmt.Scanln(&Name)
			fmt.Println("Please input task GetTime: ")
			fmt.Scanln(&GetTime)
			fmt.Println("Please input task RunTime: ")
			fmt.Scanln(&RunTime)
			fmt.Println("Please input task Size: ")
			fmt.Scanln(&Size)
			size(Size)
			fmt.Println("Please input task Tape: ")
			fmt.Scanln(&Tape)
			fmt.Println("Please input task Printer: ")
			fmt.Scanln(&Printer)

			files = append(files, file.File{
				Name:    Name,
				GetTime: GetTime,
				RunTime: RunTime,
				Size:    Size,
				Tape:    Tape,
				Printer: Printer})

			fmt.Println("input 1 to go on inputing tasks")
			fmt.Println("input 2 to enter running")
			fmt.Scan(&flag)
			fmt.Println("-----------------")
		} else if flag == 2 {
			fmt.Println("the result is:")
			hrrn.Hrrn(files)
			break
		} else {
			fmt.Println("input error, please input again")
			fmt.Scan(&flag)
			fmt.Println("-----------------")
		}
	}
}

func size(Size int) {
	for {
		if Size < 0 || Size > 64 {
			fmt.Println("the size of the main store is 64MB")
			fmt.Println("please input the task Size again:")
			fmt.Scanln(&Size)
		} else {
			return
		}
	}
}
