package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"strings"
	"wwxiong.com/go_dev/tool/qmap"
)

const (
	filePath       = "./tool/excel_to_json/file/"
	sourceFilePath = filePath + "demo.xlsx"
	destFilePath   = filePath + "demo.json"
)

var Label = []qmap.QM{
	{"_id": "100", "name": "标签", "type": "label", "parent_id": "1"},
}

func ExcelToJson() {

	f, err := excelize.OpenFile(sourceFilePath)
	if err != nil {
		log.Fatal(err)
	}

	sheetMap := f.GetSheetMap()
	sheetName, ok := sheetMap[1]
	if !ok {
		sheetName = "Sheet1"
	}
	rows, _ := f.GetRows(sheetName)

	result := make([]qmap.QM, 0)
	for i, row := range rows {
		// 跳过表头
		if i == 0 {
			continue
		}
		js := qmap.QM{}
		js["_id"] = row[0]
		js["name"] = row[1]
		js["desc"] = row[2]
		// 针对含有多个标签的处理
		label := row[3]
		ids := make([]string, 0)
		labels := strings.Split(label, "；")
		for _, s := range labels {
			l := strings.Split(s, "/")
			fmt.Println(l)
			for _, qm := range Label {
				if l[1] == qm.String("name") {
					id := qm.String("_id")
					ids = append(ids, id)
					break
				}
			}
		}
		js["label"] = ids
		js["type"] = "kd_attack"
		result = append(result, js)
	}

	b, _ := json.Marshal(result)
	// 创建一个文件并打开以进行写入
	file, err := os.Create(destFilePath)
	if err != nil {
		fmt.Println("无法创建文件:", err)
	}
	defer file.Close()

	// 将文本内容写入文件
	_, err = file.WriteString(string(b))
	if err != nil {
		fmt.Println("无法写入文件:", err)
	}
	fmt.Println(string(b))
}

func main() {
	ExcelToJson()
}
