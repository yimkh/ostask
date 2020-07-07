package hrrn

import (
	"fmt"

	file "github.com/yimkh/ostask/jobScheduling/fileInfo"
)

//Hrrn is the arithmatic
func Hrrn(files []file.File) {
	first := start(files)

	fmt.Println(first)

	otherFiles := []file.File{}

	for _, v := range files {
		if v != first {
			otherFiles = append(otherFiles, v)
		}
	}

	wait := first.RunTime + first.GetTime

	hrrn(otherFiles, wait)
}

//find start
func start(files []file.File) file.File {
	first := file.File{}

	var times []int

	for k := range files {
		times = append(times, files[k].GetTime)
	}

	firstTime := min(times)

	for _, v := range files {
		if v.GetTime == firstTime {
			first.GetTime = v.GetTime
			first.Name = v.Name
			first.RunTime = v.RunTime
			first.Tape = v.Tape
			first.Size = v.Size
			first.Printer = v.Printer
		}
	}

	return first
}

func min(l []int) (min int) {
	min = l[0]
	for _, v := range l {
		if v < min {
			min = v
		}
	}
	return
}

func hrrn(files []file.File, wait int) {
	if len(files) == 0 {
		fmt.Println("End")
		return
	}

	next := files[0]

	nexta := float32(next.RunTime + wait - next.GetTime)
	nextb := nexta / float32(next.RunTime)

	for _, v := range files {
		a := float32(v.RunTime + wait - v.GetTime)
		b := a / float32(v.RunTime)

		if b > nextb {
			next = v
			nexta = a
			nextb = b
		}
	}

	fmt.Println(next)

	otherFiles := []file.File{}

	for _, v := range files {
		if v != next {
			otherFiles = append(otherFiles, v)
		}
	}

	wait = wait + next.RunTime
	hrrn(otherFiles, wait)
}
