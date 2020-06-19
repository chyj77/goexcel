package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Luxurioust/excelize"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:rc@05926383111@tcp(39.108.0.21:3306)/dev?collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8&tls=false")
	if err != nil {
		fmt.Println(err)
	}

	//关闭数据库
	defer db.Close()

	excelFileName := "F:\\锐创\\一站式学生社区门禁数量11.27(1).xlsx"
	// sql := "insert into device_group (group_id,group_name,group_location,controll_device,devices,status,organ_id,type) values(?,?,?,?,?,0,2,0)"
	xlsx, err := excelize.OpenFile(excelFileName)

	updateSql := "update `device` set build = ? where logic_node_id =?"

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get sheet index.
	// index := xlsx.GetSheetIndex("Sheet3")
	// Get all the rows in a sheet.
	rows, _ := xlsx.GetRows("Sheet4")
	index := 0
	for _, row := range rows {
		if index == 0 {
			index = index + 1
		} else {
			cellValue0 := row[3]
			cellValue5 := row[5]
			// fmt.Println(cellValue5)
			if cellValue5 != "" {
				stmt, _ := db.Prepare(updateSql)
				//插⼊⼀⾏数据
				_, er := stmt.Exec(cellValue0, cellValue5)

				if er != nil {
					fmt.Println(er)
					os.Exit(1)
				}
				//LastInsertId返回一个数据库生成的回应命令的整数。
				//返回插入的ID
				// insID, _ := ret.LastInsertId()
				// fmt.Println(insID)
				index = index + 1
			}
		}
	}
}
