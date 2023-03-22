package main

import (
	"flag"
	"html/template"
	"log"
	"os"
)

/*
   @Auth: menah3m
   @Desc: 使用模板文件生成 SQL 语句
*/

type Param struct {
	OrgID  int
	DBName string
}

func main() {
	/* 参数1 template文件  参数2 输出文件  参数3 orgID的最大值 */

	// 1.读取命令行参数
	var templateFile, outputFile string
	var orgID int
	flag.StringVar(&templateFile, "t", "sqlt.tmpl", "sql语句的模板文件")
	flag.StringVar(&outputFile, "o", "output.sql", "解析渲染完成后的输出文件")
	flag.IntVar(&orgID, "n", 5, "最大租户ID-orgID")
	// 2.指定输出位置
	output, err := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	defer output.Close()

	if err != nil {
		log.Printf("output.txt open failed. err: %v\n", err)
	}

	// 3.解析模板
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("parse template file failed. err: %v\n", err)
	}

	// 4.渲染模板
	p := &Param{DBName: "test", OrgID: -1}
	// 租户id 为 -1 时 单独处理
	err = tmpl.Execute(output, p)
	if err != nil {
		log.Printf("template execute failed. err:%v\n", err)
	}

	for i := 0; i <= orgID; i++ {
		p.OrgID = i
		err = tmpl.Execute(output, p)
		if err != nil {
			log.Printf("template execute failed. err:%v\n", err)
		}
	}
}
