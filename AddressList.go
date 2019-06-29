package main

import "fmt"

func main(){
	scanNum()
}

func scanNum() {
	fmt.Println("添加联系人信息，请按1")
	fmt.Println("删除联系人信息，请按2")
	fmt.Println("查询联系人信息，请按3")
	fmt.Println("编辑联系人信息，请按4")

	var num int
	fmt.Scan(&num)

	switchType(num)
}

func addPerson() {
	
}

func switchType(n int) {
	switch n {
	case 1:
		//添加联系人的操作
	case 2:
		// 删除联系人的操作
	case 3:
		// 查询联系人的操作
		case 4:
			// 编辑联系人的操作	
	}
}