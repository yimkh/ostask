package arithmatic

import (
	"fmt"
	"log"
)

//Bank is banker algorithm
func Bank(Available []int, Max [][]int, Allocation [][]int, Need [][]int, Request []int, Work []int, Finish []bool, pid int, serial []int, work [][]int) {
	for i := 0; i < len(Request); i++ {
		if Request[i] <= Need[pid][i] {
			continue
		} else {
			log.Fatalln("所需资源超过宣布最大值")
		}
	}

	for i := 0; i < len(Request); i++ {
		if Request[i] <= Available[i] {
			continue
		} else {
			log.Fatalln("尚无足够资源，需等待")
		}
	}

	for i := 0; i < len(Available); i++ {
		Available[i] = Available[i] - Request[i]
	}

	for i := 0; i < len(Allocation[0]); i++ {
		Allocation[pid][i] = Allocation[pid][i] + Request[i]
	}

	for i := 0; i < len(Need[0]); i++ {
		Need[pid][i] = Need[pid][i] - Request[i]
	}

	if Safe(Available, Max, Allocation, Need, Request, Work, Finish, pid, serial, work) {
		fmt.Println("完成分配")
	} else {
		fmt.Println("系统处于不安全状态，等待")
	}

	fmt.Println("Max:", Max)
	fmt.Println("Allocation:", Allocation)
	fmt.Println("Need:", Need)
	fmt.Println("Available:", Available)
	fmt.Println("Work:", work)
	fmt.Println("安全序列：", serial)
}

//Safe is to check safe
func Safe(Available []int, Max [][]int, Allocation [][]int, Need [][]int, Request []int, Work []int, Finish []bool, pid int, serial []int, work [][]int) bool {
	Work = Available
	for i := 0; i < len(Max); i++ {
		Finish = append(Finish, false)
	}

	temp := 0
step2:
	for i := temp; i < len(Finish); i++ {
		if Finish[i] == false {
			for j := 0; j < len(Work); j++ {
				if Need[i][j] <= Work[j] {
					if j == len(Work)-1 {
						serial = append(serial, i)
						work = append(work, Work)

						for j := 0; j < len(Work); j++ {
							Work[j] = Work[j] + Allocation[i][j]
						}
						Finish[i] = true
						if i == len(Finish)-1 {
							temp = 0
						} else {
							temp = i + 1
						}
						goto step2
					}
				} else {
					break
				}
			}
		} else {
			continue
		}

		if i == len(Finish)-1 {
			return false
		}
	}

	return true
}
