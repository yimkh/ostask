package main

import (
	"fmt"
	"log"
	"os"
)

var (
	//Available is available
	Available []int
	//Max is max
	Max [][]int
	//Allocation is allocation
	Allocation [][]int
	//Need is need
	Need [][]int
	//Request is request
	Request []int
	//Work is work
	Work []int
	//Finish is finish
	Finish []bool
	pid    int   //请求资源进程号
	serial []int //用于存储安全序列
	work   [][]int
)

//初始化
func init() {

	//资源数量为10,5,7

	//资源数量为10,5,7

}

//Bank is
func Bank() {
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

	if Safe() {
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

//Safe is safe check
func Safe() bool {
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

func main() {
lable:
	var sel int
	fmt.Println("********************银行家算法********************")
	fmt.Println()
	fmt.Println("\t1、测试数据一")
	fmt.Println("\t2、测试数据二")
	fmt.Println("\t3、测试数据三")
	fmt.Println("\t4、退出")
	fmt.Println()
	fmt.Println("************************************************")
	fmt.Print("请做出你的选择：")
	fmt.Scanln(&sel)
	switch sel {
	//资源总数量为10,5,7
	case 1:
		//测试数据1
		Available = []int{3, 3, 2}
		Max = [][]int{{7, 5, 3}, {3, 2, 2}, {9, 0, 2}, {2, 2, 2}, {4, 3, 3}}
		Allocation = [][]int{{0, 1, 0}, {2, 0, 0}, {3, 0, 2}, {2, 1, 1}, {0, 0, 2}}
		Need = [][]int{{7, 4, 3}, {1, 2, 2}, {6, 0, 0}, {0, 1, 1}, {4, 3, 1}}
		Request = []int{1, 0, 2}
		pid = 1
	case 2:
		//测试数据2
		Available = []int{2, 3, 0}
		Max = [][]int{{7, 5, 3}, {3, 2, 2}, {9, 0, 2}, {2, 2, 2}, {4, 3, 3}}
		Allocation = [][]int{{0, 1, 0}, {3, 0, 2}, {3, 0, 2}, {2, 1, 1}, {0, 0, 2}}
		Need = [][]int{{7, 4, 3}, {0, 2, 0}, {6, 0, 0}, {0, 1, 1}, {4, 3, 1}}
		Request = []int{3, 3, 0}
		pid = 4
	case 3:
		//测试数据3
		Available = []int{2, 3, 0}
		Max = [][]int{{7, 5, 3}, {3, 2, 2}, {9, 0, 2}, {2, 2, 2}, {4, 3, 3}}
		Allocation = [][]int{{0, 1, 0}, {3, 0, 2}, {3, 0, 2}, {2, 1, 1}, {0, 0, 2}}
		Need = [][]int{{7, 4, 3}, {0, 2, 0}, {6, 0, 0}, {0, 1, 1}, {4, 3, 1}}
		Request = []int{0, 2, 0}
		pid = 0
	case 4:
		os.Exit(0)
	default:
		fmt.Println("没有此选项")
	}

	//执行银行家算法
	Bank()
	goto lable
}
