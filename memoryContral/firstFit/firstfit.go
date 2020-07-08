package firstfit

import (
	module "github.com/yimkh/ostask/memoryContral/module"
	"fmt"
)

//Add is to add a new task
func Add(memoryList []module.Block, taskSize int, taskName string) []module.Block {
	flag := 0

	for k, v := range memoryList {
		if v.State == 0 {
			freeSize := v.End - v.Start
			if freeSize > taskSize {
				memoryList = append(memoryList, module.Block{
					Name: taskName,
					Start: v.Start,
					End: v.Start + taskSize,
					State: 1,
				})
	
				//use this to update element in slice
				memoryList[k] = module.Block{
					Name: "",
					Start: v.Start + taskSize,
					End: v.End,
					State: 0,
				}

				flag = 1
				break
			} else if freeSize == taskSize {
				memoryList = append(memoryList, module.Block{
					Name: taskName,
					Start: v.Start,
					End: v.Start + taskSize,
					State: 1,
				})
	
				memoryList = memoryList[:k+copy(memoryList[k:], memoryList[k+1:])]

				flag = 1
				break
			} else {
				continue
			}
		}	
	}

	if flag == 0 {
		fmt.Println("NOT ENTER")
	} else {
		fmt.Println("SUCCESSFULLY ENTER")
	}

	sort(memoryList)

	return memoryList
}

//Delete is to delete a block
func Delete(memoryList []module.Block, taskName string) []module.Block {
	flag := 0
	for k := 0; k < len(memoryList); k++ {
		//find
		if memoryList[k].Name == taskName {
			memoryList[k] = module.Block{
				Name: "",
				Start: memoryList[k].Start,
				End: memoryList[k].End,
				State: 0,
			}	

			flag = 1
		}
	}

	//combine free block
	label:
	for k := 0; k < len(memoryList)-1; k++ {
		if memoryList[k].State == 0 && memoryList[k+1].State == 0 {
			memoryList[k+1] = module.Block{
				Name: "",
				Start: memoryList[k].Start,
				End: memoryList[k+1].End,
				State: 0,
			}

			// m++
			memoryList = append(memoryList[:k], memoryList[k+1:]...)
			goto label
		}
	}

	//not find
	if flag == 0 {
		fmt.Println("NOT EXIST")
	} else {
		fmt.Println("DELETE SUCCECCFULLY")
	}

	return memoryList
}

func sort(memoryList []module.Block) {
	flag := module.Block{}

	for k := 0; k < len(memoryList); k++ {
		if k != len(memoryList)-1 {
			if memoryList[k].Start > memoryList[k+1].Start {
				flag = memoryList[k+1]

				m := k
				for ; memoryList[m].Start > flag.Start; m-- {
					memoryList[m+1] = memoryList[m] 
				}
				memoryList[m+1] = flag

			}
		}
	}
}