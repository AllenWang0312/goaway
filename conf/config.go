package conf

var env = debug

var proto = "http://"
//const RediaHost  ="127.0.0.1"
//const DBHost  ="127.0.0.1"
//var FSRoot = "/home/wpc/products/netimg/file"
//var ProjectRoot = "/home/wpc/products/netimg"

const RediaHost  ="localhost"
//const RediaHost  ="127.0.0.1"
const DBHost  ="122.51.205.29"

var FSRoot = "F:/workspace/work"
var ProjectRoot = FSRoot + "/goaway"



//var FSRoot = "Z:/photos/"

var HOST = "localhost"

// var HOST = if(env == debug){"192.168.0.100"}else if (env == test){"10.10.29.249"}else{"122.51.205.29"}
var PORT = "8080"
var SERVER = proto + HOST + ":" + PORT

var FILE_PORT = "8081"
var FILE_SERVER = proto + HOST + ":" + FILE_PORT
