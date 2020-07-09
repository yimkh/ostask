package main

import (
	"fmt"

	control "github.com/yimkh/ostask/indepDevice/control"
	"github.com/yimkh/ostask/indepDevice/device"
)

func main() {
	types := [4]device.Type{}
	types[0].Name = "A"
	types[0].Count = 3
	types[0].Remain = 3
	types[0].Address = 0

	types[1].Name = "B"
	types[1].Count = 2
	types[1].Remain = 2
	types[1].Address = 3

	types[2].Name = "C"
	types[2].Count = 4
	types[2].Remain = 4
	types[2].Address = 5

	types[3].Name = "D"
	types[3].Count = 1
	types[3].Remain = 1
	types[3].Address = 9

	devices := [10]device.Device{}
	for i := 0; i < 3; i++ {
		devices[i].Number = i
		devices[i].Status = 1
		devices[i].Remain = 0
		devices[i].Jobname = ""
		devices[i].Inumber = i
	}

	for i := 3; i < 5; i++ {
		devices[i].Number = i
		devices[i].Status = 1
		devices[i].Remain = 0
		devices[i].Jobname = ""
		devices[i].Inumber = i
	}

	for i := 5; i < 9; i++ {
		devices[i].Number = i
		devices[i].Status = 1
		devices[i].Remain = 0
		devices[i].Jobname = ""
		devices[i].Inumber = i
	}

	for i := 9; i < 10; i++ {
		devices[i].Number = i
		devices[i].Status = 1
		devices[i].Remain = 0
		devices[i].Jobname = ""
		devices[i].Inumber = i
	}

	fmt.Println(types)
	fmt.Println(devices)

	types, devices = control.Allocation(types, devices, "taskA", "A", 1)
	fmt.Println(types)
	fmt.Println(devices)

	types, devices = control.Recycle(types, devices, "taskA", "A")
	fmt.Println(types)
	fmt.Println(devices)
}
