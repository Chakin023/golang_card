package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person
	// // update the value in side the struct
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// contactInfo
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@mail.com",
			zipCode: 12345,
		},
	}

	//1.1 add jimPointer and &
	// jimPointer := &jim //ex: 00x23ab
	// jimPointer.updateName("Jimmy")
	// jim.print()

	// pointer shortcut เมื่อเขียน function ถูก ไม่จำเป็นต้องใส่ & ก็ได้ GO จะรู้เรื่องโดยอัตโนมัติถึง type ที่ส่งไป
	jim.updateName("ChangeChange")
	jim.print()
}

//Receiver function for struct
func (p person) print() {
	fmt.Printf("%+v", p)
}

//1.0 Must use the pointer
//1.2 add * to receiver function and change to (*pointerToPerson)
//*person = type address ใน RAM
//*pointerToPerson = value ที่อยู่ใน address ของ RAM
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	//value ที่ -> (*00x23ab).firstName = "ค่าใหม่"
}
