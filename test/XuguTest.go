package main

import (
	"database/sql"
	"fmt"
	_ "github.com/zhengweiye/xugu-driver-go"
)

// linux：/usr/lib64/libxugusql.so
// windows：C:\Windows\System32\xugusql.dll
func main() {
	dbObj, err := sql.Open("xugusql", "IP=192.168.0.239;DB=brain;User=brain_sys;PWD=brain_sys@123;Port=5138;AUTO_COMMIT=on;CHAR_SET=UTF8")
	if err != nil {
		panic(err)
	}
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>%+v\n", dbObj)
	rows, err := dbObj.Query("select * from sys_role", []any{}...)
	if err != nil {
		panic(err)
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", rows)
}
