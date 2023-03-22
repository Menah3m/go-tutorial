package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
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

	var orgID int
	flag.IntVar(&orgID, "n", 5, "最大租户ID-orgID")
	// 2.指定输出位置

	var output bytes.Buffer

	// 3.解析模板
	const text = "\n--{{if (eq .OrgID -1)}}模板{{else}}{{.OrgID}}{{end}}租户分界线 该文本由工具自动生成\n\nINSERT INTO ecp_exam{{if (ne .OrgID -1)}}_{{.OrgID}}{{end}}.`ac_funcresource`(`idx`, `funccode`, `restype`, `respath`, `compackname`, `resname`, `status`, `creater`, `createrName`, `createTime`, `updator`, `updatorName`, `updateTime`, `orgName`, `orgCode`, `resMethod`)\nselect '1591961189558964226', idx, 'pageBtn', NULL, 'unlockStudentAccount', '账号解锁', 0, 'lyaiedu2021', '超级管理员', '2022-11-14 09:08:19', NULL, NULL, '2022-11-14 09:08:19', NULL, '{{.OrgID}}', NULL\nfrom ac_function where funcname = '班级管理';\ncommit;\n\nINSERT INTO ecp_exam{{if (ne .OrgID -1)}}_{{.OrgID}}{{end}}.`ac_funcresource`(`idx`, `funccode`, `restype`, `respath`, `compackname`, `resname`, `status`, `creater`, `createrName`, `createTime`, `updator`, `updatorName`, `updateTime`, `orgName`, `orgCode`, `resMethod`)\nselect '1591961189558964226', idx, 'pageBtn', NULL, 'unlockStudentAccount', '账号解锁', 0, 'lyaiedu2021', '超级管理员', '2022-11-14 09:08:19', NULL, NULL, '2022-11-14 09:08:19', NULL, '{{.OrgID}}', NULL\nfrom ac_function where funcname = '班级管理';\ncommit;\n\nINSERT INTO ecp_exam{{if (ne .OrgID -1)}}_{{.OrgID}}{{end}}.`ac_funcresource`(`idx`, `funccode`, `restype`, `respath`, `compackname`, `resname`, `status`, `creater`, `createrName`, `createTime`, `updator`, `updatorName`, `updateTime`, `orgName`, `orgCode`, `resMethod`)\nselect '1591961189558964226', idx, 'pageBtn', NULL, 'unlockStudentAccount', '账号解锁', 0, 'lyaiedu2021', '超级管理员', '2022-11-14 09:08:19', NULL, NULL, '2022-11-14 09:08:19', NULL, '{{.OrgID}}', NULL\nfrom ac_function where funcname = '班级管理';\ncommit;\n\n\n"

	tmpl, err := template.New("").Parse(text)
	if err != nil {
		log.Printf("parse template file failed. err: %v\n", err)
	}

	// 4.渲染模板
	p := &Param{DBName: "test", OrgID: -1}
	// 租户id 为 -1 时 单独处理
	err = tmpl.Execute(&output, p)
	if err != nil {
		log.Printf("template execute failed. err:%v\n", err)
	}

	for i := 0; i <= orgID; i++ {
		p.OrgID = i
		err = tmpl.Execute(&output, p)
		if err != nil {
			log.Printf("template execute failed. err:%v\n", err)
		}
	}

	fmt.Println(output.String())
}
