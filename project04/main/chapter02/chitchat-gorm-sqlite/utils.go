package main

import (
	"./data"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// 配置类
type Configuration struct {
	Address      string		// 初始化的服务器地址和端口
	ReadTimeout  int64		// 读数据超时
	WriteTimeout int64		// 写数据超时
	Static       string		// 静态文件配置目录
}

var config Configuration
var logger *log.Logger

// Convenience function for printing to stdout
func p(a ...interface{}) {
	fmt.Println(a)
}

func init() {
	loadConfig()
	// 设置日志记录的文件路径
	file, err := os.OpenFile("E:/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	// 初始化日志对象
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

// 读取json配置文件
func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

// Convenience function to redirect to the error message page
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// Checks if the user is logged in and has a session, if not err is not nil
func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = data.Session{UUID: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
		info("check success : " + sess.UUID + "===" + sess.Email)
	}
	return
}

// 1.使用 New() 在创建时就为其添加一个模板名称，并且执行 t.Execute() 会默认去寻找该名称进行数据融合；
// 2.使用 ParseFiles() 创建模板可以一次指定多个文件加载多个模板进来，但是就不可以使用 t.Execute() 来执行数据融合；
//   但是 ParseFiles() 可以通过 ExecuteTemplate 来进行数据融合，因为该函数可以指定模板名，
//   因此，实例模板就可以知道要去加载自己内部的哪一个模板进行数据融合。

// 当然无论使用 New() 还是 ParseFiles() 创建模板，都是可以使用 ExecuteTemplate() 来进行数据融合，
// 但是对于 Execute() 一般与 New() 创建的模板进行配合使用。

// parse HTML templates
// pass in a list of file names, and get a template
// 首先先记住一个原则 template.New() 和 ParseFiles() 最好不要一起使用，
// 如果非要一起使用，那么要记住，New(“TName”) 中的 TName 必须要和 header.tmpl 中定义的{{define name}}中的 name 同名。
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// 但是正常的做法应该是这样的，同样的 ExecuteTemplate() 中输入的 name 也必须和模板中 {{define name}} 相同。
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	// ParseFiles 函数对模版文件进行语法分析，并创建出相应的模版。
	// 为了捕捉语法分析过程中可能会产生的错误，程序使用了Must函数去包围ParseFiles函数的执行结果
	// 这样当ParseFiles函数返回错误的时候，Must函数就会向用户返回相应的错误报告。
	templates := template.Must(template.ParseFiles(files...))
	// 模版内的内容与传递的数据进行合并，生成HTML
	templates.ExecuteTemplate(writer, "layout", data)
}

// for logging
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// version
func version() string {
	return "0.1"
}
