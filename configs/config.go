package conf

var env = debug

var proto = "http://"

//var FSRoot = "Z:/photos/"
var FSRoot = "F:/workspace/work/goaway"

//var FSRoot = "/home/wpc/files/netimg"

var HOST = "localhost"

// var HOST = if(env == debug){"192.168.0.100"}else if (env == test){"10.10.29.249"}else{"122.51.205.29"}
var PORT = "8080"
var SERVER = proto + HOST + ":" + PORT

var FILE_PORT = "8081"
var FILE_SERVER = proto + HOST + ":" + FILE_PORT
