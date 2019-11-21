package main

import (
	"../../../conf"
	"fmt"
	"net/http"
)

func main() {
	//mux := http.NewServeMux()
	////获取当前路径
	////wd, err := os.Getwd()
	////if err != nil {
	////} else {
	////	fmt.Println(wd)
	////}
	////func StripPrefix(prefix string, h Handler) Handler
	//// 给定url 删除前缀
	//mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(conf.FSRoot+"/meituri"))))
	//mux.Handle("/file",http.StripPrefix("/",http.FileServer(http.Dir(conf.FSRoot+"/file"))))
	////mux.Handle("/mzitu", http.StripPrefix("/mzitu", http.FileServer(http.Dir(data+"mzitu"))))
	////mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(data+"jp"))))
	////mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(data+"cn"))))
	//http.ListenAndServe(":8081", mux)
	http.Handle("/", http.FileServer(http.Dir(conf.FSRoot)))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
