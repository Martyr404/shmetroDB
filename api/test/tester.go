package main

import (
	"fmt"
	"shmetroDB/orm"
)

func main() {
	res, err := orm.ParseCarriageNumber("jY010288")
	fmt.Println(res, err)
}
