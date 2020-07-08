package main

import (
	"fmt"

	firstfit "github.com/yimkh/ostask/memoryContral/firstFit"
	"github.com/yimkh/ostask/memoryContral/module"
)

func main() {
	memoryList := []module.Block{}

	memoryList = append(memoryList, module.Block{
		Name:  "os",
		Start: 0,
		End:   5,
		State: 1,
	})

	memoryList = append(memoryList, module.Block{
		Name:  "taskfirst",
		Start: 5,
		End:   10,
		State: 1,
	})

	memoryList = append(memoryList, module.Block{
		Name:  "taskthird",
		Start: 10,
		End:   14,
		State: 1,
	})

	memoryList = append(memoryList, module.Block{
		Name:  "",
		Start: 14,
		End:   26,
		State: 0,
	})

	memoryList = append(memoryList, module.Block{
		Name:  "tasksecond",
		Start: 26,
		End:   32,
		State: 1,
	})

	memoryList = append(memoryList, module.Block{
		Name:  "",
		Start: 32,
		End:   128,
		State: 0,
	})

	memoryList = firstfit.Add(memoryList, 6, "taskforth")
	fmt.Println(memoryList)
	memoryList = firstfit.Delete(memoryList, "taskthird")
	fmt.Println(memoryList)
	memoryList = firstfit.Delete(memoryList, "tasksecond")
	fmt.Println(memoryList)

}
