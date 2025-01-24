package main

import (
	"database/sql"
	"fmt"
	_ "github.com/zhengweiye/xugu-driver-go"
)

func main() {
	dbObj, err := sql.Open("xugusql", "IP=10.254.47.239;DB=brain;User=brain_sys;PWD=brain_sys@123;Port=5138;AUTO_COMMIT=on;CHAR_SET=UTF8")
	if err != nil {
		panic(err)
	}
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>%+v\n", dbObj)
	rows, err := dbObj.Query("select * from sys_user", []any{})
	if err != nil {
		panic(err)
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", rows)
}
