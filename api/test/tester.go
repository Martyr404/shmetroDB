package main

import (
	"fmt"
	"shmetroDB/orm"
)

func main() {
	res, err := orm.ParseCarriageNumber("012350")
	fmt.Println(res, err)
}
