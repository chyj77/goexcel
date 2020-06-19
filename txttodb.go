package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type ZWave struct {
	Type     string `json:"type"`
	Qid      string `json:"qid"`
	Deviceid string `json:"deviceid"`
}

func main() {
	err := HandleText("F:\\锐创\\gateway.txt")
	if err != nil {
		panic(err)
	}
}

func HandleText(textfile string) error {
	file, err := os.Open(textfile)
	if err != nil {
		log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
		return err
	}
	defer file.Close()

	db, err := sql.Open("mysql", "root:rc@05926383111@tcp(39.108.0.21:3306)/dev?collation=utf8mb4_unicode_ci&charset=utf8mb4,utf8&tls=false")
	if err != nil {
		fmt.Println(err)
	}

	//关闭数据库
	defer db.Close()

	var zwave ZWave

	insertSql := "INSERT INTO `isurpass_gateway` (`type`, `qid`, `deviceid`, `device_name`, `organ_id`) VALUES (?, ?, ?, 'Z-Wave网关', 2)"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// line := scanner.Text() // or
		line := scanner.Bytes()

		//do_your_function(line)
		// fmt.Printf("%s\n", line)

		json.Unmarshal(line, &zwave)

		stmt, _ := db.Prepare(insertSql)
		_, err := stmt.Exec(zwave.Type, zwave.Qid, zwave.Deviceid)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
		return err
	}

	return nil
}
