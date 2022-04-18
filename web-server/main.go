package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "----- hello ")
}

func login(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	fmt.Fprintf(w, fmt.Sprintln("have no name"))
	// }
	// contents, err := ioutil.ReadFile(`.\w\x8yexx`)
	// if err != nil {
	// 	fmt.Fprintf(w, err.Error())
	// 	return
	// }
	fmt.Fprintf(w, `<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8">
	<title>菜鸟教程(runoob.com)</title>
	</head>
	<body>
		<h1>我的第一个标题</h1>
		<p>我的第一个段落。</p>
	</body>
	</html>`)
}

func main() {
	http.HandleFunc("/w", login)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
