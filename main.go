package main

import (
	"net/http"
	"log"
	"fmt"
	"html/template"
	"strconv"
	"regexp"
)

func sayHelloName(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() // 解析参数，默认不会解析
	fmt.Fprintf(w, "Hello cocoyo!") // 写入w输出到客户端
}

func login(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() // 解析参数，默认不会解析
	if r.Method == "GET" {
		t,_ := template.ParseFiles("./views/login.gtpl")
		t.Execute(w, nil)
	} else {
		// 验证表单的输入
		// 必填字段
		if len(r.Form.Get("username")) == 0 {
			// 为空处理
		}
		// 数字
		getint,err := strconv.Atoi(r.Form.Get("age"))

		if err != nil {
			// 数字转化出错了，那么可能就不是数字
		}
		if getint > 100 {
			// 太大了
		}
		// 正则匹配
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
			// 不是数字处理
		}
		// 多选需要 r.Form["interest"] 获取  不能使用Get
		fmt.Fprintf(w, "提交的表单参数是:username: " + r.Form.Get("username"))
		fmt.Fprintf(w, "提交的表单参数是:password: " + r.Form.Get("password"))
		fmt.Fprintf(w, "提交的表单参数是:age: " + r.Form.Get("age"))
		fmt.Fprintf(w, "提交的表单参数是:interest: " + r.Form["interest"])
	}
}

func main()  {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	// 监听端口和启动服务器
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
