package crud

import (
	"shmetroDB/orm"
	"shmetroDB/psql"
	"strconv"
)

type Error = orm.Error
type TrainInfo = orm.TrainInfo

func QueryInfo(train_id string, t *TrainInfo) *Error {
	if len(train_id) == 5 {
		line, pk := train_id[:2], train_id[2:]
		line_int_tmp, _ := strconv.Atoi(line)
		line_pop_zero := strconv.Itoa(line_int_tmp)
		sql := "select * from line" + line_pop_zero + " where pk=" + PopFrontZero(pk)

		err := psql.Init()
		if err != nil {
			e := &Error{Code: "0008", Msg: "PostgreSQL initialize error", Verbose: err}
			return e
		}
		db := psql.GetDB()
		// 执行查询
		res, err := db.Query(sql)
		if err != nil {
			e := &Error{Code: "0009", Msg: "SQL Query error", Verbose: err}
			return e
		}
		defer res.Close()
		//fmt.Printf("pk\ttrain_id\ttrain_type\ttrain_detail\n")
		for res.Next() {
			err := res.Scan(&t.Pk, &t.TrainId, &t.Train_type, &t.TrainDetail)
			if err != nil {
				e := &Error{Code: "0010", Msg: "Scan data from psql goes wrong", Verbose: err}
				return e
			}
		}
		return nil
	} else {
		e := &Error{Code: "0011", Msg: "Invalid input train_id"}
		return e
	}
}
func PopFrontZero(str string) string {
	for string(str[0]) == "0" {
		str = str[1:]
	}
	return str
}

// func main() {
// 	t := TrainInfo{TrainId: "07002"}
// 	e := QueryInfo(t.TrainId, &t)
// 	fmt.Println(t, e)
// }
