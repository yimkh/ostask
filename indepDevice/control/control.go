package control

import (
	"fmt"

	"github.com/yimkh/ostask/indepDevice/device"
)

//Allocation is to allocate device for task
func Allocation(types [4]device.Type, devices [10]device.Device, taskName string, typeName string, deviceInum int) ([4]device.Type, [10]device.Device) {
	flag := 4

	for i := 0; i < len(types); i++ {
		if types[i].Name == typeName && types[i].Remain != 0 {
			if devices[types[i].Address+deviceInum].Status == 1 && devices[types[i].Address+deviceInum].Remain == 0 {
				devices[types[i].Address+deviceInum].Jobname = taskName
				devices[types[i].Address+deviceInum].Remain = 1
				types[i].Remain--

				fmt.Println("response:")
				fmt.Printf("number is %d;  ", devices[types[i].Address+deviceInum].Number)
				fmt.Printf("inumber is %d\n", devices[types[i].Address+deviceInum].Inumber)
				return types, devices
			}
			fmt.Println("device is used or bad")
		} else {
			flag--
			if flag == 0 {
				fmt.Println("this type not exist or the number of remain is 0")
			}
		}
	}
	return types, devices
}

//Recycle is to Recycle device from task
func Recycle(types [4]device.Type, devices [10]device.Device, taskName string, typeName string) ([4]device.Type, [10]device.Device) {
	flag := 4

	for i := 0; i < len(types); i++ {
		if types[i].Name == typeName {
			flagtwo := types[i].Count

			for j := 0; j < types[i].Count; j++ {
				if devices[j+types[i].Address].Jobname == taskName {
					devices[j+types[i].Address].Jobname = ""
					devices[j+types[i].Address].Remain = 0
					types[i].Remain++
					fmt.Println("response:")
					fmt.Println("Recycle ok")

					return types, devices
				}

				//not find this task
				flagtwo--
				if flagtwo == 0 {
					fmt.Println("this task is not exist")
				}
			}
		}

		//not find this type
		flag--
		if flag == 0 {
			fmt.Println("this type is not exist")
		}
	}
	return types, devices
}
