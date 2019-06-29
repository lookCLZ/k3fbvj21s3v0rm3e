package main

import "fmt"

type Person struct {
	userName     string
	addressPhone map[string]string
}

var personList []Person

func addPerson() {
	var name string
	var address string
	var phone string
	var exit string
	var addressPhone map[string]string
	fmt.Println("请输入姓名")
	fmt.Scan(&name)
	for {
		fmt.Println("请输入电话类型")
		fmt.Scan(&address)
		fmt.Println("请输入电话号码")
		fmt.Scan(&phone)

		addressPhone[address] = phone
		fmt.Println("如果结束电话的录入，请按Q")
		fmt.Scan(&exit)
		if exit == "Q" {
			break
		} else {
			continue
		}
	}
	personList = append(personList, Person{userName: name, addressPhone: addressPhone})
	fmt.Println(personList)
}

func main() {
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

func switchType(n int) {
	switch n {
	case 1:
		//添加联系人的操作
		addPerson()
	case 2:
		// 删除联系人的操作
	case 3:
	// 查询联系人的操作
	case 4:
		// 编辑联系人的操作
	}
}
