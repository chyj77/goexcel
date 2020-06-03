package main

import (
	"fmt"
	"os"

	"github.com/Luxurioust/excelize"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type DeviceGroup struct {
	Id             int    `orm:"pk;auto;"`
	GroupId        int    `orm:"column(group_id)"`
	GroupName      string `orm:"column(group_name)"`
	GroupLocation  string `orm:"column(group_location)"`
	ControllDevice string `orm:"column(controll_device)"`
	Devices        string `orm:"column(devices)"`
	Status         int    `orm:"column(status)"`
	OrganId        int    `orm:"column(organ_id)"`
	Type           int    `orm:"column(type)"`
}

func main() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:rc@05926383111@tcp(39.108.0.21:3306)/dev?collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8&tls=false")
	//注册model
	orm.RegisterModel(new(DeviceGroup))
	orm.Debug = true
	var myOrmer orm.Ormer
	myOrmer = orm.NewOrm()

	excelFileName := "F:\\锐创\\一站式学生社区门禁数量11.27(1).xlsx"
	// sql := "insert into device_group (group_id,group_name,group_location,controll_device,devices,status,organ_id,type) values(?,?,?,?,?,0,2,0)"
	xlsx, err := excelize.OpenFile(excelFileName)

	var deviceGroup DeviceGroup

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get sheet index.
	// index := xlsx.GetSheetIndex("Sheet3")
	// Get all the rows in a sheet.
	rows, _ := xlsx.GetRows("Sheet3")
	index := 0
	for _, row := range rows {
		if index == 0 {
			index = index + 1
		} else {
			size := len(row)
			cellValue3 := row[3]
			cellValue5 := row[5]
			var cellValue8 string
			if size == 9 {
				cellValue8 = row[8]
			}
			if cellValue5 != "" {
				deviceGroup.GroupId = index
				deviceGroup.GroupName = cellValue3
				deviceGroup.GroupLocation = cellValue3
				deviceGroup.ControllDevice = cellValue5
				deviceGroup.Devices = cellValue8
				deviceGroup.Status = 0
				deviceGroup.OrganId = 2
				deviceGroup.Type = 0
				myOrmer.Insert(&deviceGroup)
				index = index + 1
			}
		}
	}
}
