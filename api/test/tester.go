package main

import (
	"fmt"
	"shmetroDB/orm"
)

func main() {
	res, err := orm.ParseCarriageNumber("jC40093")
	fmt.Println(res, err)
}
