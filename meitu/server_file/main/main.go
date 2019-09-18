package main

import (
	"net/http"
)

//var dataDir="Z:/photos/"
var data = "I:/wpc/projects/golang/"

func main() {
	mux := http.NewServeMux()
	//获取当前路径
	//wd, err := os.Getwd()
	//if err != nil {
	//} else {
	//	fmt.Println(wd)
	//}
	//func StripPrefix(prefix string, h Handler) Handler
	// 给定url 删除前缀
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(data+"images"))))
	//mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(data+"jp"))))
	//mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(data+"cn"))))
	http.ListenAndServe(":8081", mux)
}
