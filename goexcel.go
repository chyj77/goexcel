package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Luxurioust/excelize"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type EbxRookBox struct {
	Id          int    `orm:"pk;auto;"`
	Addrs       string `json:"addrs" orm:"column(addrs)"`
	ProjectCode string `json:"projectCode" orm:"column(projectCode)"`
	Mac         string `json:"mac" orm:"column(mac)"`
}

type Device struct {
	DeviceId    string  `json:"deviceId"`
	Lines       []Lines `json:"lines"`
	ProjectCode string  `json:"projectCode"`
	Type        string  `json:"type"`
}

type Lines struct {
	Addr string `json:"addr"`
}

func main() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:rc@05926383111@tcp(39.108.0.21:3306)/dev?collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8&tls=false")
	//注册model
	orm.RegisterModel(new(EbxRookBox))
	// orm.Debug = true
	var myOrmer orm.Ormer
	myOrmer = orm.NewOrm()
	var ebxrookboxs EbxRookBox

	excelFileName := "F:\\锐创\\一站式学生社区门禁数量11.27(1).xlsx"
	sql := "SELECT 0 id, GROUP_CONCAT(a.addr) addrs,a.project_code projectCode,a.mac FROM `ebx_rook_box_channels` a where  a.mac= ? and a.project_code= ? "
	xlsx, err := excelize.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get sheet index.
	// index := xlsx.GetSheetIndex("Sheet3")
	// Get all the rows in a sheet.
	rows, _ := xlsx.GetRows("Sheet4")
	index := 1
	for _, row := range rows {
		// cellValue5 := row[5]
		cellValue6 := row[6]
		cellValue7 := row[7]

		var array []Device
		// println("cellValue6=", cellValue6)
		macs := strings.Split(cellValue6, ",")
		var mString string
		var cell string
		for _, mac := range macs {
			myOrmer.Raw(sql, mac, cellValue7).QueryRow(&ebxrookboxs)
			// fmt.Print(ebxrookboxs, "\t")
			if ebxrookboxs.Addrs != "" {
				addrs := ebxrookboxs.Addrs
				addr := strings.Split(addrs, ",")
				var addrArray []Lines
				var lines Lines
				for _, adr := range addr {
					lines.Addr = adr
					addrArray = append(addrArray, lines)
					// fmt.Println(addrArray)
				}
				var device Device
				device.DeviceId = mac
				device.Lines = addrArray
				device.ProjectCode = cellValue7
				device.Type = "ebxrook"
				array = append(array, device)
				mjson, _ := json.Marshal(array)
				mString = string(mjson)
				cell = "I" + strconv.Itoa(index)
			}
		}
		// println(mString)
		xlsx.SetCellValue("Sheet4", cell, mString)
		index = index + 1
	}

	if err := xlsx.SaveAs(excelFileName); err != nil {
		println(err.Error())
	}
}
