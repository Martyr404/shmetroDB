// package crud
package main

import (
	"fmt"
	"shmetroDB/orm"
	"shmetroDB/psql"
	"strconv"
)

type Error = orm.Error

type Train_Backend struct {
	pk              int
	train_id        string
	train_type      string
	train_detail    string
	Carriage_number []string
	Carriage_type   []string
}

func QueryInfo(train_id string) (*Train_Backend, *Error) {
	if len(train_id) == 5 {
		line, pk := train_id[:2], train_id[2:]
		line_int_tmp, _ := strconv.Atoi(line)
		line_pop_zero := strconv.Itoa(line_int_tmp)
		sql := "select * from line" + line_pop_zero + " where pk=" + pk

		err := psql.Init()
		if err != nil {
			e := &Error{Code: "0008", Msg: "PostgreSQL initialize error", Verbose: err}
			return nil, e
		}
		db := psql.GetDB()
		// 执行查询
		res, err := db.Query(sql)
		if err != nil {
			e := &Error{Code: "0009", Msg: "SQL Query error", Verbose: err}
			return nil, e
		}
		defer res.Close()
		t := Train_Backend{}
		//fmt.Printf("pk\ttrain_id\ttrain_type\ttrain_detail\n")
		for res.Next() {
			err := res.Scan(&t.pk, &t.train_id, &t.train_type, &t.train_detail)
			if err != nil {
				e := &Error{Code: "0010", Msg: "Scan data from psql goes wrong", Verbose: err}
				return nil, e
			}

			// fmt.Printf(
			// 	"%d\t%s\t\t%s\t\t%s \n",
			// 	t.pk, t.train_id, t.train_type, t.train_detail,
			// )
		}
		return &t, nil
	} else {
		e := &Error{Code: "0011", Msg: "Invalid input train_id"}
		return nil, e
	}
}
func main() {
	t, _ := QueryInfo("07052")
	fmt.Println(t)
}
