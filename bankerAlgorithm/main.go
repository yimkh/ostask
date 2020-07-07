package main

import (
	"fmt"
	"os"

	"github.com/yimkh/ostask/bankerAlgorithm/arithmatic"
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

//init is to inititial all resource
func init() {
	//资源数量为10,5,7
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
	arithmatic.Bank(Available, Max, Allocation, Need, Request, Work, Finish, pid, serial, work)
	goto lable
}
