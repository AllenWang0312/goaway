package conf


const Env = Pre
const proto = "http://"
const GormDebug=true

//release
// const FSHost="122.51.205.29"
// const RediaHost = "127.0.0.1"
// const DBHost = "127.0.0.1"

// const Muri="/muri"
// var FSRoot = "/home/wpc/products/netimg"
// var ProjectRoot = "/home/wpc/products/netimg"

//work
const FSHost="127.0.0.1"
const RediaHost = "122.51.205.29"
const DBHost = "122.51.205.29"
var Muri="/muri"
var FSRoot = "F:/work"
var ProjectRoot = FSRoot + "/goaway"
//home
//const FSHost="127.0.0.1"
//const RediaHost = "122.51.205.29"
//const DBHost = "122.51.205.29"
//const Muri="/meituri_cn_like"
////const Muri="/muri"
//var FSRoot = "Z:/photos"
//var ProjectRoot ="E:/work/goaway"


const RediaPass  ="qunsi003"
const MysqlPass  ="qunsi003"

const HOST = "localhost"
const PORT = "8080"
const FILE_PORT = "8081"
// var HOST = if(env == debug){"192.168.0.100"}else if (env == test){"10.10.29.249"}else{"122.51.205.29"}
var FSMuri = FSRoot+Muri

var SERVER = proto + HOST + ":" + PORT

var FILE_SERVER = proto + FSHost + ":" + FILE_PORT
