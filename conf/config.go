package conf

var env = debug

var proto = "http://"

//const RediaHost = "127.0.0.1"
//const DBHost = "127.0.0.1"
//
//var FSRoot = "/home/wpc/products/netimg"
//var ProjectRoot = "/home/wpc/products/netimg"

const RediaHost = "122.51.205.29"
const DBHost = "122.51.205.29"

var FSRoot = "F:/workspace/work"
var ProjectRoot = FSRoot + "/goaway"

//const RediaHost  ="122.51.205.29"
//const DBHost  ="122.51.205.29"
//var FSRoot = "Z:/photos/"
//var ProjectRoot ="E:/work/goaway"
//

var HOST = "localhost"

// var HOST = if(env == debug){"192.168.0.100"}else if (env == test){"10.10.29.249"}else{"122.51.205.29"}
var PORT = "8080"
var SERVER = proto + HOST + ":" + PORT

var FILE_PORT = "8081"
var FILE_SERVER = proto + "122.51.205.29" + ":" + FILE_PORT
