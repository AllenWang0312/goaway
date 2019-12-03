package conf

const Env = Pre
//release
//const RediaHost = "127.0.0.1"
//const DBHost = "127.0.0.1"
//var FSRoot = "/home/wpc/products/netimg"
//var FSMuri = FSRoot+"/muri"
//var ProjectRoot = "/home/wpc/products/netimg"
//work
const RediaHost = "122.51.205.29"
const DBHost = "122.51.205.29"
var FSRoot = "F:/workspace/work"
var FSMuri = FSRoot+"/muri"
var ProjectRoot = FSRoot + "/goaway"
//home
//const RediaHost  ="122.51.205.29"
//const DBHost  ="122.51.205.29"
//var FSRoot = "Z:/photos"
//var FSMuri = FSRoot+"/meituri_cn"
//var ProjectRoot ="E:/work/goaway"
//
const proto = "http://"

const GormDebug=false
const RediaPass  ="qunsi003"
const MysqlPass  ="qunsi003"

const HOST = "localhost"
const PORT = "8080"
const FILE_PORT = "8081"
// var HOST = if(env == debug){"192.168.0.100"}else if (env == test){"10.10.29.249"}else{"122.51.205.29"}

var SERVER = proto + HOST + ":" + PORT
var FILE_SERVER = proto + "122.51.205.29" + ":" + FILE_PORT
