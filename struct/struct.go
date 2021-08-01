package main

import (

	"fmt"
)
type Address struct {
	HNo int
	street string

}

type customer struct {
	custName   string
	custID     int
	course    string
	Address
}

func main() {

	initialise()
}	
func objreference(cust *customer) {
	fmt.Println("customer name:", (*cust).custName, c.custID, c.course, c.Address)
}
func valueref(cust customer)  {
	fmt.Println("customer data", cust.custName)
}
func initialise()  {
	cust1 := customer {}
	cust1.custName = "John"
	cust1.custID = 2021
	cust1.course = "Golang"
	cust1.HNo = 5
	cust1.street = "Bnglr"
	    fmt.Println(cust1)
	    fmt.Println()
		add := Address {
		HNo : 7,
		street : "Hyd",
  }

    cust2 := customer 
	custName == "Stephen"
	custID = 2022
	course = "Java"
	Address = add
	   
	}
	fmt.Println(cust2)
	objreference(&cust2)
	valueref(cust2)
	data := make(map[int]customer)
	data[10] = cust1
	data[11] = cust1
	data[12] = cust2
	fmt.Println(data)
	//Add
	var add1 = data[10]
	fmt.Println("customer info on key 10", add)
	//Update
	data[11] = cust1
	fmt.Println("update key 11", data)
	//Delete
	delete(data, 12)
	fmt.Println("delete key 12", data)
}
