package main

// import (
// 	"fmt"
// 	"shmetroDB/orm"
// 	"shmetroDB/psql"
// )

// type trainInfo = orm.TrainInfo

// // 应对pgsql表名大小写区分问题，pgsql默认全转化小写识别，需要加"vAlue"区分大小写，sql语句要用\"转义符号表示双引号
// func main() {
// 	sql := "select * from " + "\"lineAL\"" + " where pk=" + "1"

// 	err := psql.Init()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	db := psql.GetDB()
// 	// 执行查询
// 	res, err := db.Query(sql)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer res.Close()
// 	for res.Next() {
// 		t := trainInfo{}
// 		err = res.Scan(&t.Pk, &t.TrainId, &t.Train_type, &t.TrainDetail)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Println(t)
// 	}
// }
